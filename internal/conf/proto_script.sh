#!/bin/bash

kratos proto client --proto_path=. --proto_path=$GOPATH/src internal/conf/config.v1.proto
