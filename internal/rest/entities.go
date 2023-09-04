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
	ID       int    `json:"id,omitempty"`
	Document string `json:"document,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
}

type Response struct {
	Message string `json:"message"`
}

type ProductResponse struct {
	ID          int     `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Category    string  `json:"category,omitempty"`
	Status      int     `json:"status,omitempty"`
	ImagePath   string  `json:"image_path,omitempty"`
}

type ProductRequest struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Status      int     `json:"status"`
	ImagePath   string  `json:"image_path"`
}

type OrderRequest struct {
	CustomerID int         `json:"customer_id"`
	PaymentID  int         `json:"payment_id"`
	OrderItems []OrderItem `json:"products"`
}

type OrderResponse struct {
	ID           int               `json:"id,omitempty"`
	CustomerID   int               `json:"customer_id,omitempty"`
	Status       int               `json:"status,omitempty"`
	CreatedAt    *time.Time        `json:"created_at,omitempty"`
	UpdatedAt    *time.Time        `json:"updated_at,omitempty"`
	Products     []ProductResponse `json:"products,omitempty"`
	*PaymentRest `json:"payment,omitempty"`
}

type OrderItem struct {
	Product  ProductRequest `json:"product"`
	Quantity int64          `json:"quantity"`
}

type PaymentRest struct {
	ID          int        `json:"id"`
	PaymentType int        `json:"payment_type"`
	CreatedAt   *time.Time `json:"created_at"`
}

type PaymentCallback struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
}
