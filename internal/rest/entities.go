package rest

import "time"

type CustomerRequest struct {
	Document string `json:"document"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type CustomerResponse struct {
	ID       int    `json:"id"`
	Document string `json:"document"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type Response struct {
	Message string `json:"message"`
}

type ProductResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Status      int     `json:"status"`
	ImagePath   string  `json:"imagePath"`
}
type ProductRequest struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Status      int     `json:"status"`
	ImagePath   string  `json:"imagePath"`
}

type OrderRequest struct {
	CustomerID int         `json:"customer_id"`
	PaymentID  int         `json:"payment_id"`
	Status     int         `json:"status"`
	CreatedAt  *time.Time  `json:"created_at"`
	UpdatedAt  *time.Time  `json:"updated_at"`
	OrderItems []OrderItem `json:"products"`
}

type OrderItem struct {
	Product  ProductRequest `json:"product"`
	Quantity int64          `json:"quantity"`
}

type Payment struct {
	ID          int
	PaymentType int
	CreatedAt   *time.Time
}
