#!/bin/bash

echo "start to generate"

protoc --go_out=./protopb ./proto/hello/hello.proto
protoc --go-grpc_out=./protopb ./proto/hello/hello.proto

echo "done"

