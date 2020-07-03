#!bin/bash

# destroyer service
protoc proto/destroyer-service.proto --go_out=plugins=grpc:src/destroyer-service

# deathstar service
protoc proto/deathstar-service.proto --go_out=plugins=grpc:src/deathstar-service