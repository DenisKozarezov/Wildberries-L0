package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

const (
	StreamName = "ORDERS"
	Subjects   = "ORDERS.*"
)

func PublishNewOrder(js jetstream.JetStream, i int) {
	js.PublishAsync("ORDERS.new", []byte("hello message "+strconv.Itoa(i)))
	log.Printf("[JetStream] Published a new order %d.\n", i)
}

func ConnectToNATS() (jetstream.JetStream, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		return nil, fmt.Errorf("Could not connect to NATS! %w", err)
	} else {
		log.Println("Connected to NATS!")
	}

	// Create jetstream context from nats connection
	js, err := CreateJetStream(ctx, nc)

	if err != nil {
		return nil, fmt.Errorf("Could not create a JetStream! %w", err)
	} else {
		log.Println("Stream created!")
	}

	return js, nil
}

func CreateJetStream(ctx context.Context, nc *nats.Conn) (jetstream.JetStream, error) {
	js, _ := jetstream.New(nc)
	log.Println("JetStream context created.")

	// Create a stream
	_, err := js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     StreamName,
		Subjects: []string{Subjects},
	})

	if err != nil {
		return nil, err
	}

	return js, nil
}

func main() {
	js, err := ConnectToNATS()

	if err != nil {
		log.Fatalln(err)
	}

	// Publish some messages
	for i := 0; i < 100; i++ {
		PublishNewOrder(js, i)
	}
}
