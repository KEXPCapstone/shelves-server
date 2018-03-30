#!/usr/bin/env bash
set -e
echo "Building Linux executable for library microservice..."
GOOS=linux go build 
echo "Building docker cotainer for library microservice..."
docker build -t kexpcapstone/library .
echo "Pushing container to DockerHub"
docker push kexpcapstone/library
go clean