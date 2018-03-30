#!/usr/bin/env bash
export ADDR=localhost:4000
export TLSCERT=$(GOPATH)/src/github.com/KEXPCapstone/shelves-server/tls/fullchain.pem
export TLSKEY=$(GOPATH)/src/github.com/KEXPCapstone/shelves-server/tls/privkey.pem
export REDISADDR=localhost:6379
export SESSIONKEY=password
export DBADDR=localhost:27017
export LIBRARYSVCADDR=localhost:4001
export SHELVESSVCADDR=localhost:4002
go install && gateway
