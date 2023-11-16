package database

import (
	"encoding/json"
	"fmt"
)

type Item struct {
	Chrt_id      int
	Track_number string
	Price        int
	Rid          string
	Name         string
	Sale         int
	Size         string
	Total_price  int
	Nm_id        int
	Brand        string
	Status       int
}

type Payment struct {
	Transaction   string
	Request_id    string
	Currency      string
	Provider      string
	Amount        int
	Payment_dt    int
	Bank          string
	Delivery_cost int
	Goods_total   int
	Custom_fee    int
}

type Delivery struct {
	Name    string
	Phone   string
	Zip     string
	City    string
	Address string
	Region  string
	Email   string
}

type Order struct {
	Order_uid          string
	Track_number       string
	Entry              string
	Delivery           Delivery
	Payment            Payment
	Items              []*Item
	Locale             string
	Internal_signature string
	Customer_id        string
	Delivery_service   string
	Shardkey           string
	Sm_id              int
	Date_created       string
	Oof_shard          string
}

type OrderRepository struct {
}

func (self *OrderRepository) SelectAll() ([]Order, error) {
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
		err = json.Unmarshal(([]byte)(jsonData), &order)

		if err != nil {
			return nil, fmt.Errorf("Unable to unmarshal json from the database. %w", err)
		}

		orders = append(orders, order)
	}

	return orders, err
}

func (self *OrderRepository) SelectByUID(uid string) (*Order, error) {
	row := db.QueryRow("SELECT DISTINCT * FROM orders WHERE order_uid =$1", uid)

	var orderUid string
	var jsonData string
	err := row.Scan(&orderUid, &jsonData)

	if err != nil {
		return nil, fmt.Errorf("Unable to scan row: %w", err)
	}

	var order Order
	err = json.Unmarshal(([]byte)(jsonData), &order)

	if err != nil {
		return nil, fmt.Errorf("Unable to unmarshal json from the database. %w", err)
	}

	return &order, err
}

func (self *OrderRepository) Insert(order *Order) error {
	return nil
}
