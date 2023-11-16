package services

import (
	"log"
	db "myapp/database"
)

var rep db.IDatabaseRepository[db.Order] = &db.OrderRepository{}

func SelectAllOrders() ([]db.Order, error) {
	orders, err := rep.SelectAll()

	if err != nil {
		log.Fatalf("Unable to query orders! Reason: %s\n", err)
		return nil, err
	}

	log.Printf("[SELECT ALL ORDERS]: %s", orders)
	log.Printf("Selected rows count: %d\n", len(orders))

	return orders, err
}

func SelectOrderByUID(uid string) (*db.Order, error) {
	order, err := rep.SelectByUID(uid)

	if err != nil {
		log.Fatalf("Unable to query an order! Reason: %s\n", err)
		return nil, err
	}

	log.Printf("[SELECT ORDER BY UID]: %s", uid)
	log.Printf("Selected row: %s\n", *order)
	return order, err
}

func AddNewOrder(order *db.Order) error {
	err := rep.Insert(order)

	if err != nil {
		log.Fatalf("Couldn't add a new order! Reason: %s\n", err)
	}
	return err
}
