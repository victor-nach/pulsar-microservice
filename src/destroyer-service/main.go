package main

import (
	"fmt"
	"github.com/victor-nach/pulsar-microservice/src/destroyer-service/server"
	"log"
)

func main() {
	fmt.Println("Starting destroyer service ...")
	if err := server.Run(); err != nil {
		log.Fatalf("Error starting service : %v", err)
	}

}
