#!/bin/bash

kratos proto client --proto_path=. --proto_path=$GOPATH/src api/common/v1/auth.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/common/v1/config.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/common/v1/env.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/common/v1/error.v1.proto
