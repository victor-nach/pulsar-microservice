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
	//
	//url, ok := os.LookupEnv("SERVICE_URL")
	//if !ok {
	//	url = "pulsar://localhost:6650"
	//}
	//client, err := pulsar.NewClient(pulsar.ClientOptions{
	//	URL:               url,
	//	OperationTimeout:  30 * time.Second,
	//	ConnectionTimeout: 30 * time.Second,
	//})
	//
	//if err != nil {
	//	log.Fatalf("Could not instantiate Pulsar client: %v", err)
	//}
	//
	//defer client.Close()
	//
	//
	//serviceName, ok := os.LookupEnv("SERVICE_NAME")
	//if !ok {
	//	serviceName = "deathstar"
	//}
	//topic, ok := os.LookupEnv("TOPIC_NAME")
	//if !ok {
	//	topic = "targets.acquired"
	//}
	//consumer, err := client.Subscribe(pulsar.ConsumerOptions{
	//	Topic:            topic,
	//	SubscriptionName: serviceName,
	//	Type:             pulsar.Shared,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer consumer.Close()
	//
	//for i := 0; i < 10; i++ {
	//	msg, err := consumer.Receive(context.Background())
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	var target Target
	//	err = json.Unmarshal(msg.Payload(), &target)
	//	if err != nil {
	//		fmt.Println("could not unmarshall.")
	//	}
	//	fmt.Println(target)
	//
	//	consumer.Ack(msg)
	//}
	//
	//if err := consumer.Unsubscribe(); err != nil {
	//
	//	log.Fatal(err )
	//}
}

