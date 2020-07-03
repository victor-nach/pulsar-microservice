package handler

import (
	"context"
	"fmt"
	"log"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/victor-nach/pulsar-microservice/src/deathstar-service/server/repo"
)

type EventService struct {
	Consumer	pulsar.Consumer
	Repo		repo.RepoInterface
}

func (e *EventService) Run () {
	log.Println("Listeneing for events")
	for i := 0; i < 10; i++ {
		msg, err := e.Consumer.Receive(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		err = e.Repo.SaveTarget(msg)
		if err != nil {
			fmt.Errorf("Error saving target: %v", err)
			return
		}
		// Acknowledge message
	}
}
