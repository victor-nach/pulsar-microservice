#!bin/bash

# greet service
protoc ${PWD}/../proto/deathstar-service.proto --go_out=plugins=grpc:src/deathstar-service

# calculator service
#protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.