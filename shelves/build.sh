#!/usr/bin/env bash
set -e
echo "Building Linux executable for shelves microservice..."
GOOS=linux go build 
echo "Building docker cotainer for shelves microservice..."
docker build -t kexpcapstone/shelves .
echo "Pushing container to DockerHub"
docker push kexpcapstone/shelves
go clean