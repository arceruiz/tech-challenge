package port

import "tech-challenge/internal/canonical"

type CustomerService interface {
	Create(canonical.Customer) (*canonical.Customer, error)
	Login(user canonical.Customer) (string, error)
	Bypass() (string, error)
}

type OrderService interface {
	GetOrders() ([]canonical.Order, error)
	CreateOrder(order canonical.Order) (canonical.Order, error)
	UpdateOrder(id string, updatedOrder canonical.Order) (canonical.Order, error)
	GetByID(id string) (canonical.Order, error)
	GetByStatus(id string) ([]canonical.Order, error)
}

type ProductService interface {
	GetProducts() ([]canonical.Product, error)
	CreateProduct(product canonical.Product) error
	UpdateProduct(id string, updatedProduct canonical.Product) error
	GetByID(id string) (*canonical.Product, error)
	GetByCategory(id string) ([]canonical.Product, error)
	Remove(id string) error
}
