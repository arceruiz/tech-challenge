package entities

import "time"

type Order struct {
	ID         int64     `json:"id"`
	CustomerID int64     `json:"customerId"`
	PaymentID  int64     `json:"paymentId"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Total      float64   `json:"total"`
}
