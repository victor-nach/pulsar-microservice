# pulsar-microservice

[![Maintainability](https://api.codeclimate.com/v1/badges/6b59f00c3add4598f0a0/maintainability)](https://codeclimate.com/github/victor-nach/pulsar-microservice/maintainability) [![Test Coverage](https://api.codeclimate.com/v1/badges/6b59f00c3add4598f0a0/test_coverage)](https://codeclimate.com/github/victor-nach/pulsar-microservice/test_coverage) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A simple micro-service application that communicate through a pub-sub messaging system (Apache Pulsar).)

## How to run

### Locally
- Clone this project
- run `make start` to build docker images and start all the services

## Run services seperately
ALternatively you can run the services seperately by
- first running the dependencies (mongo & pulsar) `make run-dep`
- running the deathstar service independently `make run-deathstar`
- running the destroyer service independently `make run-destroyer`
- running the api gateway service independently `make run-api`


### Project Management
The pivotal tracker board for the project can be found here - https://www.pivotaltracker.com/n/projects/2455810
