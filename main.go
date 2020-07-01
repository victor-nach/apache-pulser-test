package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

func main() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://localhost:6650",
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})

	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}

	fmt.Println(client)
	defer client.Close()


	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "topic-1",
		SubscriptionName: "my-sub",
		Type:             pulsar.Shared,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	for i := 0; i < 10; i++ {
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Received message msgId: %#v -- content: '%s'\n",
			msg.ID(), string(msg.Payload()))

		consumer.Ack(msg)
	}

	    if err := consumer.Unsubscribe(); err != nil {

		log.Fatal(err )
	}

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "my-topic",
	})

	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte("hello"),
	})

	defer producer.Close()

	if err != nil {
		fmt.Println("Failed to publish message", err)
	}
	fmt.Println("Published message")
}
