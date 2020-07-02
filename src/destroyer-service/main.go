package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

type Target struct {
	Id			string	`json:"id"`
	Message 	string	`json:"message"`
	CreatedOn	time.Time
}
func main() {
	fmt.Println("Starting destroyer service ...")

	url, ok := os.LookupEnv("PULSAR_URL")
	if !ok {
		url = "pulsar://localhost:6650"
	}
	fmt.Println(url)

	// create a new pulsar client
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:        url,
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})

	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}

	defer client.Close()

	serviceName, ok := os.LookupEnv("SERVICE_NAME")
	if !ok {
		serviceName = "destroyer"
	}
	topic, ok := os.LookupEnv("TOPIC_NAME")
	if !ok {
		topic = "targets.acquired"
	}
	// create a new producer instance
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
		Name: serviceName,
	})

	payload, _ := json.Marshal(Target{
		Id: "12345",
		Message: "hello from this side",
	})

	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: payload,
	})

	defer producer.Close()

	if err != nil {
		fmt.Println("Failed to publish message", err)
	}
	fmt.Println("Published message")
}

