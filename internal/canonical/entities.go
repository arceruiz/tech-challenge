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
	*Payment
	Status     OrderStatus
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	Total      float64
	OrderItems []OrderItem //orderProduct
}

type OrderStatus int
type PaymentStatus int

const (
	ORDER_CANCELLED  OrderStatus = 0
	ORDER_RECEIVED   OrderStatus = 1
	ORDER_PREPARING  OrderStatus = 2
	ORDER_READY      OrderStatus = 3
	ORDER_DELIEVERED OrderStatus = 4
)

const (
	PAYMENT_INIT  PaymentStatus = 0
	PAYMENT_OK    PaymentStatus = 1
	PAYMENT_ERROR PaymentStatus = 2
	PAYMENT_NOK   PaymentStatus = 2
)

var MapPaymentStatus = map[string]PaymentStatus{
	"OK":        PAYMENT_OK,
	"NOK":       PAYMENT_NOK,
	"ERROR":     PAYMENT_ERROR,
	"INIT":      PAYMENT_INIT,
	"":          PAYMENT_ERROR,
	"COMPLETED": PAYMENT_OK,
	"PENDING":   PAYMENT_INIT,
}

type OrderItem struct {
	Product
	Quantity int64
}

type Payment struct {
	ID          int
	PaymentType int
	CreatedAt   *time.Time
	Status      PaymentStatus
}
