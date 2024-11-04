package models

type Order struct {
	ID           int    `json:"id"`
	CustomerName string `json:"customer_name"`
	OrderDate    string `json:"order_date"`
	Status       string `json:"status"`
}
