package models

type Order struct {
    ID           int     `json:"id"`
    CustomerName string  `json:"customer_name"`
    Total        float64 `json:"total"`
}
