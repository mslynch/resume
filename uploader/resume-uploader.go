package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

type AuthInfo struct {
	ApiUrl             string `json:"apiUrl"`
	AuthorizationToken string `json:"authorizationToken"`
}

type UploadUrlRequest struct {
	BucketId string `json:"bucketId"`
}

type UploadInfo struct {
	UploadUrl          string `json:"uploadUrl"`
	AuthorizationToken string `json:"authorizationToken"`
}

func main() {
	applicationKeyId := getenv("RESUME_APPLICATION_KEY_ID")
	applicationKey := getenv("RESUME_APPLICATION_KEY")
	bucketId := getenv("RESUME_BUCKET_ID")
	filename := getenv("RESUME_FILENAME")

	httpClient := http.Client{
		Timeout: 60 * time.Second,
	}

	fmt.Println("Authenticating...")
	authInfo := authenticate(&httpClient, applicationKeyId, applicationKey)
	fmt.Println("Determining upload URL...")
	uploadInfo := getUploadUrl(&httpClient, &authInfo, bucketId)
	fmt.Println("Uploading file...")
	uploadFile(&httpClient, &uploadInfo, filename)
	fmt.Println("File uploaded!")
}

func getenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatal("missing required environment variable: " + key)
	}
	return value
}

func authenticate(httpClient *http.Client, applicationKeyId, applicationKey string) AuthInfo {
	req, _ := http.NewRequest("GET", "https://api.backblazeb2.com/b2api/v2/b2_authorize_account", nil)
	req.SetBasicAuth(applicationKeyId, applicationKey)
	res, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	checkResponse(res, "b2_authorize_account")

	var authInfo AuthInfo
	json.NewDecoder(res.Body).Decode(&authInfo)
	return authInfo
}

func getUploadUrl(httpClient *http.Client, authInfo *AuthInfo, bucketId string) UploadInfo {
	requestBody, _ := json.Marshal(UploadUrlRequest{BucketId: bucketId})
	requestBodyReader := bytes.NewReader(requestBody)
	req, err := http.NewRequest("POST", authInfo.ApiUrl+"/b2api/v2/b2_get_upload_url", requestBodyReader)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", authInfo.AuthorizationToken)
	res, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	checkResponse(res, "b2_get_upload_url")

	var uploadInfo UploadInfo
	json.NewDecoder(res.Body).Decode(&uploadInfo)
	return uploadInfo
}

func uploadFile(httpClient *http.Client, uploadInfo *UploadInfo, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	fileBuffer := new(bytes.Buffer)
	fileBuffer.ReadFrom(file)
	fileBytes := fileBuffer.Bytes()
	fileReader := bytes.NewReader(fileBytes)
	checksum := fmt.Sprintf("%x", sha1.Sum(fileBytes))
	req, reqError := http.NewRequest("POST", uploadInfo.UploadUrl, fileReader)
	if reqError != nil {
		log.Fatal(reqError)
	}
	req.Header.Add("Authorization", uploadInfo.AuthorizationToken)
	req.Header.Add("Content-Type", "application/pdf")
	req.Header.Add("X-Bz-Content-Sha1", checksum)
	req.Header.Add("X-Bz-File-Name", "resume.pdf")
	res, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	checkResponse(res, "b2_upload_file")
}

func checkResponse(res *http.Response, apiName string) {
	if res.StatusCode != 200 {
		dump, dumpErr := httputil.DumpResponse(res, true)
		if dumpErr != nil {
			log.Fatal(dumpErr)
		}
		fmt.Printf("received bad %s response:\n", apiName)
		fmt.Println(string(dump))
		log.Fatal()
	}
}
