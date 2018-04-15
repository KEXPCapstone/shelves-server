#!/usr/bin/env bash
export ADDR=localhost:4001
export DBADDR=localhost:27017
export RELEASEDB=library
export RELEASECOLL=releases
go install && library