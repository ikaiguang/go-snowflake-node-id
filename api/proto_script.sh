#!/bin/bash

protoc -I. -I$GOPATH/src --go_out=. --go_opt=paths=source_relative ./api/*.proto
# kratos proto client --proto_path=. --proto_path=$GOPATH/src ./api/node-id.error.v1.proto
# kratos proto client --proto_path=. --proto_path=$GOPATH/src ./api/node-id.resource.v1.proto
