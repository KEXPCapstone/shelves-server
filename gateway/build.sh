#!/usr/bin/env bash
set -e
echo "Building Linux executable for gateway..."
GOOS=linux go build 
echo "Building docker cotainer for gateway..."
docker build -t kexpcapstone/gateway .
echo "Pushing container to DockerHub"
docker push kexpcapstone/gateway
go clean
