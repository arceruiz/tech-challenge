package adapters

import (
	"database/sql"
	"tech-challenge/internal/canonical"
	"tech-challenge/internal/repository"
	"tech-challenge/internal/repository/port"
)

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepo() port.OrderRepository {
	return &orderRepository{repository.New()}
}

func (r *orderRepository) GetOrders() ([]canonical.Order, error) {
	return nil, nil
}

func (r *orderRepository) CreateOrder(order canonical.Order) (canonical.Order, error) {
	return canonical.Order{}, nil
}

func (r *orderRepository) UpdateOrder(id string, updatedOrder canonical.Order) (canonical.Order, error) {
	return canonical.Order{}, nil
}

func (r *orderRepository) GetByID(id string) (canonical.Order, error) {
	return canonical.Order{}, nil
}

func (r *orderRepository) GetByStatus(id string) ([]canonical.Order, error) {
	return nil, nil
}
