# drone-gcs

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-gcs/status.svg)](http://beta.drone.io/drone-plugins/drone-gcs)
[![Go Doc](https://godoc.org/github.com/drone-plugins/drone-gcs?status.svg)](http://godoc.org/github.com/drone-plugins/drone-gcs)
[![Go Report](https://goreportcard.com/badge/github.com/drone-plugins/drone-gcs)](https://goreportcard.com/report/github.com/drone-plugins/drone-gcs)
[![Join the chat at https://gitter.im/drone/drone](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/drone/drone)

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
docker build --rm=true -t plugins/gcs .
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
  -e AWS_ACCESS_KEY_ID=<token> \
  -e AWS_SECRET_ACCESS_KEY=<secret> \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  plugins/gcs --dry-run
```
