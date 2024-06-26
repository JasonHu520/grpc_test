#!/bin/bash
 protoc -I=. -I=../vendor --go_out=. info.proto
 protoc -I=. -I=./vendor --go-grpc_out=. test/info.proto
