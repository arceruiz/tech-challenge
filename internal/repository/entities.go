package repository

import "time"

type OrderJoinPayment struct {
	ID               int
	CustomerID       int
	PaymentID        *int
	Status           int
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
	Total            float64
	PaymentType      *int
	PaymentCreatedat *time.Time
	PaymentStatus    *int
}

type Order struct {
	ID         int
	CustomerID int
	PaymentID  *int
	Status     int
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	Total      float64
}
