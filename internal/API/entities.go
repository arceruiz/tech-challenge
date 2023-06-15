package api

import "time"

type Customer struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Document  string    `json:"document"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type Product struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Desc      string  `json:"desc"`
	Price     float64 `json:"price"`
	Category  string  `json:"category"`
	ImagePath string  `json:"imagePath"`
}

type OrderItems struct {
	ID        int64 `json:"id"`
	OrderID   int64 `json:"orderId"`
	ProductID int64 `json:"productId"`
	Quantity  int   `json:"quantity"`
}

type Order struct {
	ID         int64     `json:"id"`
	CustomerID int64     `json:"customerId"`
	PaymentID  int64     `json:"paymentId"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	Total      float64   `json:"total"`
}

type Payment struct {
	ID          int64     `json:"id"`
	PaymentType string    `json:"paymentType"`
	CreatedAt   time.Time `json:"createdAt"`
}
