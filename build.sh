#!/usr/bin/env bash

# This will not work unfortunately
# go get ./...

# You have to manually declare all your dependencies here
go get github.com/julienschmidt/httprouter

# Build your application to the bin directory
go build -o bin/application application.go
