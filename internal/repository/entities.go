package repository

import "time"

type Order struct {
	ID               int
	CustomerID       int
	PaymentID        *int
	Status           int
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
	Total            float64
	PaymentType      *int
	paymentCreatedat *time.Time
}
