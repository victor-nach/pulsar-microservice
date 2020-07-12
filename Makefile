gen-proto:
	#./gen-proto.sh
	protoc proto/destroyer-service.proto --go_out=plugins=grpc:src/destroyer-service && \
	protoc proto/deathstar-service.proto --go_out=plugins=grpc:src/deathstar-service && \
	protoc proto/destroyer-service.proto --go_out=plugins=grpc:src/api-gateway

build: gen-proto
	docker-compose build

start: 
	docker-compose up -d

run: build
	docker-compose up -d

# To start up the destroyer service independently
run-destroyer: 
	cd src/destroyer-service && go run main.go

# To start up the deathstar service independently
run-deathstar: gen-proto
	cd src/deathstar-service && go run main.go

# To start up the dependencies independently (mongodb, Apache pulsar)
run-dep:
	docker-compose -f docker-compose.dep.yml up

# To run up the api gateway client
run-api:
	cd src/api-gateway && go run main.go
