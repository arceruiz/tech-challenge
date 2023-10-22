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
	Status      BaseStatus
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

type BaseStatus int
type OrderStatus int
type PaymentStatus int

const (
	STATUS_ACTIVE BaseStatus = iota
	STATUS_INACTIVE
)

const (
	ORDER_READY OrderStatus = iota
	ORDER_PREPARING
	ORDER_RECEIVED
	ORDER_DELIEVERED
	ORDER_CONCLUDED
	ORDER_CANCELLED
)

const (
	PAYMENT_CREATED PaymentStatus = iota
	PAYMENT_PAYED
	PAYMENT_ERROR
	PAYMENT_FAILED
)

var MapBaseStatus = map[string]BaseStatus{ //ajustar chaves
	"ACTIVE":   STATUS_ACTIVE,
	"INACTIVE": STATUS_INACTIVE,
}

var MapOrderStatus = map[string]OrderStatus{ //ajustar chaves
	"OK":        ORDER_CANCELLED,
	"NOK":       ORDER_RECEIVED,
	"ERROR":     ORDER_PREPARING,
	"INIT":      ORDER_READY,
	"":          ORDER_DELIEVERED,
	"COMPLETED": ORDER_CONCLUDED,
}

var MapPaymentStatus = map[string]PaymentStatus{
	"OK":        PAYMENT_PAYED,
	"NOK":       PAYMENT_FAILED,
	"ERROR":     PAYMENT_ERROR,
	"INIT":      PAYMENT_CREATED,
	"":          PAYMENT_ERROR,
	"COMPLETED": PAYMENT_PAYED,
	"PENDING":   PAYMENT_CREATED,
}
