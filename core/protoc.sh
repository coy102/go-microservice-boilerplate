#!/bin/bash
go get -d github.com/micro/micro/v2/cmd/protoc-gen-micro@master
protoc --proto_path=${pwd} --micro_out=. --micro_opt=paths=source_relative --go_out=. --go_opt=paths=source_relative proto/*.proto
protoc --proto_path=${pwd} --micro_out=. --micro_opt=paths=source_relative --go_out=. --go_opt=paths=source_relative proto/health/health.proto

