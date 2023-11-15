package services

import (
	"log"
	db "myapp/database"
)

var orderRep db.OrderRepository = db.OrderRepository{}

func SelectAll() ([]db.Order, error) {
	orders, err := orderRep.SelectAll()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Printf("[SELECT ALL ORDERS]: %s", orders)
	log.Printf("Selected rows count: %d\n", len(orders))

	return orders, err
}

func SelectOrderByUid(uid string) (*db.Order, error) {
	order, err := orderRep.SelectByUID(uid)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Printf("[SELECT ORDER BY UID]: %s", uid)
	log.Printf("Selected row: %s\n", order)
	return order, err
}
