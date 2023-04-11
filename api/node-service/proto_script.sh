#!/bin/bash

kratos proto client --proto_path=. --proto_path=$GOPATH/src api/node-service/v1/resources/node_id.resource.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/node-service/v1/services/node_id.service.v1.proto
