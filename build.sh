#!/bin/bash
echo "Getting dependencies..."
go get -u github.com/golang/dep/cmd/dep
dep init
dep ensure
echo "Building..."
CGO_ENABLED=0 GOOS=linux go build -o websocketd
