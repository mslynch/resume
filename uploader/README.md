# Resume - uploader source

Files used to upload the resume document to Backblaze B2.

## Usage
Before running, the following environment variables must be present:
```
RESUME_BUCKET_ID           # the Backblaze B2 bucket into which to upload the document
RESUME_APPLICATION_KEY_ID  # the Backblaze B2 application key id used to access RESUME_BUCKET_ID
RESUME_APPLICATION_KEY     # the Backblaze B2 application key used to access RESUME_BUCKET_ID
RESUME_FILENAME            # the file to be uploaded
```

To run, use:
```shell
go run uploader/resume_uploader.go
```
