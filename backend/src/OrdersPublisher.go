package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

const (
	StreamName = "ORDERS"
	Subjects   = "ORDERS.*"
)

func PublishNewOrder(js jetstream.JetStream, data []byte) {
	js.PublishAsync("ORDERS.new", data)
	log.Printf("[JetStream] Published a new order: %s.\n", data)
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
	if js, err := CreateJetStream(ctx, nc); err != nil {
		return nil, fmt.Errorf("Could not create a JetStream! %w", err)
	} else {
		log.Println("Stream created!")
		return js, nil
	}
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

	for i := 1; i < 5; i++ {
		filepath := "../../data/order" + strconv.Itoa(i) + ".json"
		data, err := os.ReadFile(filepath)

		if err != nil {
			log.Fatalf("Couldn't read a file '%s'! %w", filepath, err)
		} else {
			log.Printf("File '%s' is successfully read.\n", filepath)
		}

		PublishNewOrder(js, data)
	}
}
