# drone-gcs

[![Build Status](https://drone.wyattjoh.com/api/badges/wyatt/drone-gcs/status.svg)](https://drone.wyattjoh.com/wyatt/drone-gcs)
[![Go Doc](https://godoc.org/github.com/wyattjoh/drone-gcs?status.svg)](http://godoc.org/github.com/wyattjoh/drone-gcs)
[![Go Report](https://goreportcard.com/badge/github.com/wyattjoh/drone-gcs)](https://goreportcard.com/report/github.com/wyattjoh/drone-gcs)

Drone plugin to publish files and artifacts to Google Cloud Storage. For the
usage information and a listing of the available options please take a look at
[the docs](DOCS.md).

## Build

Build the binary with the following commands:

```
go build
go test
```

## Docker

Build the docker image with the following commands:

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo
docker build --rm=true -t wyattjoh/drone-gcs .
```

Please note incorrectly building the image for the correct x64 linux and with
GCO disabled will result in an error when running the Docker image:

```
docker: Error response from daemon: Container command
'/bin/drone-gcs' not found or does not exist..
```

## Usage

Execute from the working directory:

```
docker run --rm \
  -e PLUGIN_SOURCE=<source> \
  -e PLUGIN_TARGET=<target> \
  -e PLUGIN_BUCKET=<bucket> \
  -e GOOGLE_APPLICATION_CREDENTIALS_CONTENTS=<application credentials json> \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  wyattjoh/drone-gcs --dry-run
```
