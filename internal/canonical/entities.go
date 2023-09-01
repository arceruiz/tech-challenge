package canonical

import "time"

type Customer struct {
	Id        int
	Document  string
	Name      string
	Email     string
	Password  string
	CreatedAt *time.Time
}

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Category    string
	Status      int
	ImagePath   string
}

type Order struct {
	ID         int
	CustomerID int
	PaymentID  int
	Status     int
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	Total      float64
	OrderItems []OrderItem //orderProduct
}

type OrderItem struct {
	Product
	Quantity int64
}

type Payment struct {
	ID          int
	PaymentType int
	CreatedAt   *time.Time
}
