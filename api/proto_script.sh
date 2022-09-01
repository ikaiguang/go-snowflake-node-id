#!/bin/bash

kratos proto client --proto_path=. --proto_path=$GOPATH/src ./api/node-id.resource.v1.proto
