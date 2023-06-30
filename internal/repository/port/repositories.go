package port

import "tech-challenge/internal/canonical"

type CustomerRepository interface {
	Create(canonical.Customer) error
	GetByEmail(email string) (*canonical.Customer, error)
}

type OrderRepository interface {
	GetOrders() ([]canonical.Order, error)
	CreateOrder(order canonical.Order) (canonical.Order, error)
	UpdateOrder(id string, updatedOrder canonical.Order) (canonical.Order, error)
	GetByID(id string) (canonical.Order, error)
	GetByStatus(id string) ([]canonical.Order, error)
}

type ProductRepository interface {
	GetProducts() ([]canonical.Product, error)
	CreateProduct(product canonical.Product) (canonical.Product, error)
	UpdateProduct(id string, updatedProduct canonical.Product) (canonical.Product, error)
	GetByID(id string) (canonical.Product, error)
	GetByCategory(id string) ([]canonical.Product, error)
}
