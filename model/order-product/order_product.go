package orderproduct

import "time"

type OrderProduct struct {
	ID int `json:"id"`
	Trx string `json:"trx"`
	OrderProductId int `json:"order_product_id"`
	ProductId int `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}