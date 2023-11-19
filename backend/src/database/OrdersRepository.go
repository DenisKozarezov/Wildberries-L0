package database

import (
	"encoding/json"
	"fmt"
)

type OrdersRepository struct {
}

func (self *OrdersRepository) SelectAll() ([]Order, error) {
	rows, err := db.Query("SELECT * FROM orders")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []Order{}
	for rows.Next() {
		var orderUid string
		var jsonData string
		err := rows.Scan(&orderUid, &jsonData)
		if err != nil {
			return nil, fmt.Errorf("Unable to scan row: %w", err)
		}

		var order Order
		err = json.Unmarshal([]byte(jsonData), &order)

		if err != nil {
			return nil, fmt.Errorf("Unable to unmarshal json from the database. %w", err)
		}

		orders = append(orders, order)
	}

	return orders, err
}

func (self *OrdersRepository) SelectByUID(uid string) (*Order, error) {
	row := db.QueryRow("SELECT DISTINCT * FROM orders WHERE order_uid =$1", uid)

	var orderUid string
	var jsonData string
	err := row.Scan(&orderUid, &jsonData)

	if err != nil {
		return nil, fmt.Errorf("Unable to scan row: %w", err)
	}

	var order Order
	err = json.Unmarshal([]byte(jsonData), &order)

	if err != nil {
		return nil, fmt.Errorf("Unable to unmarshal json from the database. %w", err)
	}

	return &order, err
}

func (self *OrdersRepository) Insert(order_uid string, data string) error {
	var err error
	_, err = db.Exec("INSERT INTO orders (order_uid, json) VALUES ($1, $2)", order_uid, data)

	if err != nil {
		return fmt.Errorf("Could not insert into database: %w", err)
	}

	return nil
}
