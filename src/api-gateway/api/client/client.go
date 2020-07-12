package client

import (
	"fmt"
	"log"
	"os"
	"github.com/victor-nach/pulsar-microservice/src/api-gateway/pb"
	"google.golang.org/grpc"
)

// Client - grpc client struct
type Client struct{}

// Interface - grpc client interface
type Interface interface {
	GetDestroyerClient() (pb.DestroyerServiceClient, error)
}

// Cli - Destroyer service grpc client
var Cli pb.DestroyerServiceClient

func init() {
	var client Client
	err := client.GetDestroyerClient()
	if err != nil {
		log.Fatal("err", err)
	} 
}

// GetDestroyerClient - get a connection to the destroyer service
func (r *Client) GetDestroyerClient() error {
	url, ok := os.LookupEnv("DESTROYER_URL")
	if !ok {
		url = "localhost:5051"
	}
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return err
	}
	Cli = pb.NewDestroyerServiceClient(conn)
	fmt.Println("connected to destroyer service")
	return nil
}
