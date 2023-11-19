package database

import "time"

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
