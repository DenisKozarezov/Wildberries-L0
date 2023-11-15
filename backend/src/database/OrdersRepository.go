package database

import (
	"fmt"
)

type Item struct {
	chrt_id      int
	track_number string
	price        int
	rid          string
	name         string
	sale         int
	size         int
	total_price  int
	nm_id        int
	brand        string
	status       int
}

type Payment struct {
	transaction   string
	request_id    string
	currency      string
	provider      string
	amount        int
	payment_dt    int
	bank          string
	delivery_cost int
	goods_total   int
	custom_fee    int
}

type Delivery struct {
	name    string
	phone   string
	zip     string
	city    string
	address string
	region  string
	email   string
}

type Order struct {
	order_uid    string
	track_number string
	entry        string
	delivery     string
}

type OrderRepository struct {
}

func (self *OrderRepository) SelectAll() ([]Order, error) {
	rows, err := db.Query("SELECT * FROM orders")

	if err != nil {
		return nil, fmt.Errorf("Unable to query orders: %w\n", err)
	}
	defer rows.Close()

	orders := []Order{}
	for rows.Next() {
		order := Order{}
		err := rows.Scan(&order.order_uid, &order.track_number, &order.entry, &order.delivery)
		if err != nil {
			return nil, fmt.Errorf("Unable to scan row: %w", err)
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (self *OrderRepository) SelectByUID(uid string) (*Order, error) {
	row := db.QueryRow("SELECT DISTINCT * FROM orders WHERE order_uid =$1", uid)

	order := Order{}
	err := row.Scan(&order.order_uid, &order.track_number, &order.entry, &order.delivery)

	if err != nil {
		return nil, err
	}

	return &order, err
}
