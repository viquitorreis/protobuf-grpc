#!/bin/bash

# Ensure GOPATH is set
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

go get -u google.golang.org/protobuf
go get -u google.golang.org/grpc

# Install protoc-gen-validate
go install github.com/envoyproxy/protoc-gen-validate@latest
go get
