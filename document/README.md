# Resume - document source

Files used to build the resume document.

## Requirements
* Docker (or `alias docker=podman`), or the following:
* Fonts Inter and IBM Plex Sans
* `wkhtmltopdf`
* `entr` (optional, for quicker development iteration)
* `npm` to deploy to Netlify


## Building the PDF
To build, run one of the following scripts. The resulting PDF will be placed at `target/resume.pdf`. To update the placeholder phone number, set the environment variable `PHONE_NUMBER`. 

With Docker:
```sh
./bin/build-docker-full
```

Without Docker:
```sh
./bin/build
```

To watch files and build upon change, run (after having run `./bin/build-docker-full` once):
```sh
./bin/dev-docker
```

Or without Docker:
```sh
./bin/dev
```

## Deployment
This project is hosted using a Netlify function.

To build the function files, run:
```sh
./bin/build-function
```

To deploy the built files to Netlify, run:
```
npm install
./bin/deploy
```
