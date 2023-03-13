
# GO-docker-simple
This repo is meant to be use as a simple way to deploy the simpliest go app.

## Build container

docker build . -t go-docker-simple

docker run -p 8080:8080 go-docker-simple
