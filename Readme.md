#Build container

docker build . -t go-docker-simple

docker run -p 8080:8080 go-docker-simple