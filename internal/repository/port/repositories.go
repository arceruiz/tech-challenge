package port

import (
	"tech-challenge/internal/canonical"
	"tech-challenge/internal/repository"
)

type CustomerRepository interface {
	Create(canonical.Customer) error
	GetByEmail(email string) (*canonical.Customer, error)
}

type OrderRepository interface {
	GetOrders() ([]canonical.Order, error)
	CreateOrder(order canonical.Order) error
	UpdateOrder(id string, updatedOrder canonical.Order) (canonical.Order, error)
	GetByID(id string) (*canonical.Order, error)
	GetByStatus(status string) ([]canonical.Order, error)
	CheckoutOrder(orderID string, payment canonical.Payment) error
}

type OrderItemRepository interface {
	GetOrderItemsFromOrderID(orderID string) ([]repository.OrderItem, error)
}

type ProductRepository interface {
	GetProducts() ([]canonical.Product, error)
	CreateProduct(product canonical.Product) error
	UpdateProduct(id string, product canonical.Product) error
	GetByID(id string) (*canonical.Product, error)
	GetByCategory(id string) ([]canonical.Product, error)
}
