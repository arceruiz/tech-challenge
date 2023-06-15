package usecases

import (
	entities "tech-challenge/internal/Domain/Entities"
)

type ICustomerUseCase interface {
	GetCustomers() ([]entities.Customer, error)
	CreateCustomer(customer entities.Customer) (entities.Customer, error)
	GetCustomer(id string) (entities.Customer, error)
	UpdateCustomer(id string, updatedCustomer entities.Customer) (entities.Customer, error)
	DeleteCustomer(id string) error
}

type IProductUseCase interface {
	GetProducts() ([]entities.Product, error)
	CreateProduct(product entities.Product) (entities.Product, error)
	GetProduct(id string) (entities.Product, error)
	UpdateProduct(id string, updatedProduct entities.Product) (entities.Product, error)
	DeleteProduct(id string) error
}

type IOrderUseCase interface {
	GetOrders() ([]entities.Order, error)
	CreateOrder(order entities.Order) (entities.Order, error)
	GetOrder(id string) (entities.Order, error)
	UpdateOrder(id string, updatedOrder entities.Order) (entities.Order, error)
	DeleteOrder(id string) error
}

type IPaymentUseCase interface {
	GetPayments() ([]entities.Payment, error)
	CreatePayment(payment entities.Payment) (entities.Payment, error)
	GetPayment(id string) (entities.Payment, error)
	UpdatePayment(id string, updatedPayment entities.Payment) (entities.Payment, error)
	DeletePayment(id string) error
}
