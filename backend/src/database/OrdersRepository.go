package database

import (
	"encoding/json"
	"fmt"
	"time"
)

type Order struct {
	Order_uid          string   `validate:"required"`
	Track_number       string   `validate:"required"`
	Entry              string   `validate:"required"`
	Delivery           Delivery `validate:"required"`
	Payment            Payment  `validate:"required"`
	Items              []Item
	Locale             string `validate:"required"`
	Internal_signature string
	Customer_id        string    `validate:"required"`
	Delivery_service   string    `validate:"required"`
	Shardkey           string    `validate:"required"`
	Sm_id              int       `validate:"required"`
	Date_created       time.Time `validate:"required"`
	Oof_shard          string    `validate:"required"`
}

type Delivery struct {
	Name    string `validate:"required"`
	Phone   string `validate:"required"`
	Zip     string `validate:"required"`
	City    string `validate:"required"`
	Address string `validate:"required"`
	Region  string `validate:"required"`
	Email   string `validate:"required"`
}

type Payment struct {
	Transaction   string `validate:"required"`
	Request_id    string
	Currency      string `validate:"required"`
	Provider      string `validate:"required"`
	Amount        int    `validate:"required"`
	Payment_dt    int64  `validate:"required"`
	Bank          string `validate:"required"`
	Delivery_cost int    `validate:"required"`
	Goods_total   int    `validate:"required"`
	Custome_fee   int
}

type Item struct {
	Chrt_id      int    `validate:"required"`
	Track_number string `validate:"required"`
	Price        int    `validate:"required"`
	Rid          string `validate:"required"`
	Name         string `validate:"required"`
	Sale         int    `validate:"required"`
	Size         string `validate:"required"`
	Total_price  int    `validate:"required"`
	Nm_id        int    `validate:"required"`
	Brand        string `validate:"required"`
	Status       int    `validate:"required"`
}

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
