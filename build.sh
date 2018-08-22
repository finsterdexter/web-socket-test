#!/bin/bash
echo "Getting dependencies..."
go get -u github.com/golang/dep/cmd/dep
dep init
dep ensure
echo "Building..."
go build -o websocketd
chmod +x websocketd
