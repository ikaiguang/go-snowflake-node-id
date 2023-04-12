#!/bin/bash

kratos proto client --proto_path=. --proto_path=$GOPATH/src api/snowflake-service/v1/errors/snowflake.error.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/snowflake-service/v1/resources/snowflake.resource.v1.proto
kratos proto client --proto_path=. --proto_path=$GOPATH/src api/snowflake-service/v1/services/snowflake.service.v1.proto
