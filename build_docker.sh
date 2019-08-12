#!/bin/bash

docker run -v "$GOPATH":/go --rm -v "$PWD":/go/src/github.com/Huhaokun/let-it-fail -w /go/src/github.com/Huhaokun/let-it-fail -e GO111MODULE=on golang:1.12.6-alpine go build  -o lif-agent ./cmd/agent

docker build -t lif-agent -f cmd/agent/Dockerfile .
