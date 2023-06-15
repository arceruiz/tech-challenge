package entities

import "time"

type Payment struct {
	ID          int64     `json:"id"`
	PaymentType string    `json:"paymentType"`
	CreatedAt   time.Time `json:"createdAt"`
}
