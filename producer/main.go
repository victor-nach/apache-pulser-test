package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

func main() {
	// create a new pulsar client

	url := os.Getenv("SERVICE_URL")
	fmt.Println(url)
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:        url,
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})

	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}

	//fmt.Println(client)
	defer client.Close()

	// create a new producer instance
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "my-topic",
	})

	// sends a message to the event broker
	// this is a blocking call, would wait till the message is acknowledged by pulser
	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte("hello"),
	})

	defer producer.Close()

	if err != nil {
		fmt.Println("Failed to publish message", err)
	}
	fmt.Println("Published message")
}
