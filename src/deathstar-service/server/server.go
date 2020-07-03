package server

import (
	"encoding/json"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/victor-nach/pulsar-microservice/src/deathstar-service/pb"
	"github.com/victor-nach/pulsar-microservice/src/deathstar-service/server/db"
	"github.com/victor-nach/pulsar-microservice/src/deathstar-service/server/handler"
	"github.com/victor-nach/pulsar-microservice/src/deathstar-service/server/repo"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"time"
)

//Run - run the server
func Run() error {
	c, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	r := repo.Repository{Collection: c}
	comsumer, err := getComsumer()
	if err != nil {
		log.Printf("Failed to connect to broker : %v", err)
		return err
	}
	h := handler.Handler{r, producer}
	s := grpc.NewServer()
	pb.RegisterDestroyerServiceServer(s, &h)
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Printf("Failed to listen : %v", err)
		return err
	}
	fmt.Println("Serving deathstar service...")
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v\n", err)
		return err
	}
	return nil
}

// getComsumer - create a pulsar consumer
func getComsumer() error {
	url, ok := os.LookupEnv("PULSAR_URL")
	if !ok {
		url = "pulsar://localhost:6650"
	}
	fmt.Println(url)
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               url,
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})

	if err != nil {
		log.Printf("Could not instantiate Pulsar client: %v", err)
		return nil, err
	}

	// create a new producer instance
	serviceName, ok := os.LookupEnv("SERVICE_NAME")
	if !ok {
		serviceName = "deathstar"
	}
	topic, ok := os.LookupEnv("TOPIC_NAME")
	if !ok {
		topic = "targets.acquired"
	}
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: serviceName,
		Type:             pulsar.Shared,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()


}
