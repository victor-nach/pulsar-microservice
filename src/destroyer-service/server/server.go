package server

import (
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/victor-nach/pulsar-microservice/src/destroyer-service/pb"
	//"github.com/victor-nach/pulsar-microservice/src/destroyer-service/server/db"
	"github.com/victor-nach/pulsar-microservice/src/destroyer-service/server/handler"
	"github.com/victor-nach/pulsar-microservice/src/destroyer-service/server/repo"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"time"
)

//Run - run the server
func Run() error {
	//collection, err := db.Connect()
	//if err != nil {
	//	log.Fatal(err)
	//}
	r := repo.Repository{}
	producer, err := getProducer()
	if err != nil {
		log.Printf("Failed to connect to broker : %v", err)
		return err
	}
	h := handler.Handler{r, producer}
	s := grpc.NewServer()
	pb.RegisterDestroyerServiceServer(s, &h)
	lis, err := net.Listen("tcp", "0.0.0.0:50051") // this is the port binding
	if err != nil {
		log.Printf("Failed to listen : %v", err)
		return err
	}
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v\n", err)
		return err
	}
	fmt.Println("Serving destroyer service...")
	return nil
}

// getProducer - create a pulsar producer
func getProducer() (pulsar.Producer, error) {
	// create a new pulsar client
	url, ok := os.LookupEnv("SERVICE_URL")
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

	//fmt.Println(client)
	defer client.Close()

	// create a new producer instance
	serviceName, ok := os.LookupEnv("SERVICE_NAME")
	if !ok {
		serviceName = "destroyer"
	}
	topic, ok := os.LookupEnv("TOPIC_NAME")
	if !ok {
		topic = "targets.acquired"
	}
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
		Name:  serviceName,
	})
	return producer, nil
}