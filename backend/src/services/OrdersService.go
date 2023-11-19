package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	db "myapp/database"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

var rep db.IDatabaseRepository[db.Order] = &db.OrdersRepository{}
var ordersCache Cache[db.Order]

func validate(order *db.Order) validator.ValidationErrors {
	validate := validator.New()
	if err := validate.Struct(order); err != nil {
		errs := err.(validator.ValidationErrors)
		for _, fieldErr := range errs {
			fmt.Printf("Field %s: %s\n", fieldErr.Field(), fieldErr.Tag())
		}
		return errs
	}
	return nil
}

func ordersHandler(msg jetstream.Msg) {
	msg.Ack()

	bytes := msg.Data()
	data := string(bytes)

	log.Printf("Received a JetStream message via callback: %s\n", data)

	var order db.Order
	if err := json.Unmarshal(bytes, &order); err != nil {
		log.Println("Unable to unmarshal from NATS!", err)
		return
	}

	if errs := validate(&order); len(errs) == 0 {
		log.Println("Validation succeeded!")

		if err := AddNewOrder(order.Order_uid, data); err != nil {
			log.Println(err)
		}
	} else {
		log.Println("Validation failed!")
	}
}

func ConnectToNATS() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	if nc, err := nats.Connect(nats.DefaultURL); err == nil {
		log.Println("Connected to NATS!")

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
	} else {
		log.Println("Could not connect to NATS!")
		log.Fatalln(err)
	}
}

func RestoreCache() {
	log.Println("Restoring the cache...")

	ordersCache.New()

	orders, err := SelectAllOrders()

	if err != nil {
		panic(fmt.Errorf("Orders cache cannot be restored! %w", err))
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
		log.Printf("Unable to query orders! %s\n", err)
		return nil, err
	}

	log.Println("[SELECT ALL ORDERS]:\n", orders)
	log.Printf("Selected rows count: %d\n", len(orders))

	return orders, err
}

func SelectOrderByUID(uid string) (*db.Order, error) {
	order, found := ordersCache.Get(uid)

	if found {
		log.Printf("Order UID = [%s] is found in cache.", uid)
		log.Printf("Found order: %s", *order)
		return order, nil
	} else {
		log.Printf("Order UID = [%s] is NOT found in cache.", uid)

		order, err := rep.SelectByUID(uid)

		if err != nil {
			fmt.Errorf("Unable to query an order UID = [%s]! %w\n", uid, err)
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
		return fmt.Errorf("Unable to add a new order! %w", err)
	}

	log.Println("New order was added into database.")

	return err
}
