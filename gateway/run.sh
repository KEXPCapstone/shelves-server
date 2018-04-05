#!/usr/bin/env bash
export ADDR=localhost:4000
export TLSCERT=${GOPATH}/src/github.com/KEXPCapstone/shelves-server/tls/fullchain.pem
export TLSKEY=${GOPATH}/src/github.com/KEXPCapstone/shelves-server/tls/privkey.pem
export REDISADDR=localhost:6379
export SESSIONKEY=password
export DBADDR=localhost:27017
export LIBRARYSVCADDR=localhost:4001
export SHELVESSVCADDR=localhost:4002
docker rm -f redissvr
docker rm -f mongodb
docker -d -p 6379:6379 --name redissvr redis
docker -d -p 27017:27017 --name mongodb mongo
go install && gateway
