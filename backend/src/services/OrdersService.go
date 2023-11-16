package services

import (
	"fmt"
	"log"
	db "myapp/database"
)

var rep db.IDatabaseRepository[db.Order] = &db.OrderRepository{}
var ordersCache Cache[db.Order]

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
}

func SelectAllOrders() ([]db.Order, error) {
	orders, err := rep.SelectAll()

	if err != nil {
		log.Printf("Unable to query orders! Reason: %s\n", err)
		return nil, err
	}

	log.Printf("[SELECT ALL ORDERS]:\n%s", orders)
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

func AddNewOrder(order *db.Order) error {
	err := rep.Insert(order)

	if err != nil {
		log.Printf("Couldn't add a new order! Reason: %s\n", err)
	}
	return err
}
