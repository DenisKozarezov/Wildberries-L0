package services

import (
	"context"
	"fmt"
	"log"
	db "myapp/database"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

var rep db.IDatabaseRepository[db.Order] = &db.OrdersRepository{}
var ordersCache Cache[db.Order]

func ordersHandler(msg jetstream.Msg) {
	msg.Ack()
	fmt.Printf("Received a JetStream message via callback: %s\n", string(msg.Data()))
}

func ConnectToNATS() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Println("Could not connect to NATS!")
		log.Fatalln(err)
	} else {
		log.Println("Connected to NATS!")
	}

	js, _ := jetstream.New(nc)
	log.Println("JetStream context created.")

	stream, err := js.Stream(ctx, "ORDERS")

	if err != nil {
		log.Printf("Could not get a stream from NATS connection! %s", err)
	}

	// Create durable consumer
	c, _ := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable:   "CONS",
		AckPolicy: jetstream.AckExplicitPolicy,
	})

	// Receive messages continuously in a callback
	c.Consume(ordersHandler)
	// defer cons.Stop()
}

func RestoreCache() {
	log.Println("Restoring the cache...")

	ordersCache.New()

	orders, err := SelectAllOrders()

	if err != nil {
		panic(fmt.Errorf("Orders cache cannot be restored! Reason: %s", err))
	}

	for i, order := range orders {
		log.Printf("[%d] Adding %s in cache...", i, order)
		ordersCache.Add(order.Order_uid, &order)
	}

	log.Println("Cache restored.")
}

func SelectAllOrders() ([]db.Order, error) {
	orders, err := rep.SelectAll()

	if err != nil {
		log.Printf("Unable to query orders! Reason: %s\n", err)
		return nil, err
	}

	log.Println("[SELECT ALL ORDERS]:\n", orders)
	log.Printf("Selected rows count: %d\n", len(orders))

	return orders, err
}

func SelectOrderByUID(uid string) (*db.Order, error) {
	order, found := ordersCache.Get(uid)

	if found {
		log.Printf("Order UID = [%s] is found in cache. Found order: %s", uid, *order)
		return order, nil
	} else {
		log.Printf("Order UID = [%s] is NOT found in cache.", uid)

		order, err := rep.SelectByUID(uid)

		if err != nil {
			log.Printf("Unable to query an order UID = [%s]! Reason: %s\n", uid, err)
			return nil, err
		}

		log.Printf("[SELECT ORDER BY UID]: %s", uid)
		log.Printf("Selected row: %s\n", *order)
		return order, err
	}
}

func AddNewOrder(order_uid string, data string) error {
	err := rep.Insert(order_uid, data)

	if err != nil {
		log.Printf("Unable to add a new order! Reason: %s\n", err)
	}
	return err
}
