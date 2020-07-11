package main

import (
	"fmt"
	"log"
	"github.com/victor-nach/pulsar-microservice/src/deathstar-service/server"
)

func main() {
	fmt.Println("Starting deathstar service ...")
	if err := server.Run(); err != nil {
		log.Fatalf("Error starting service : %v", err)
	}
}

