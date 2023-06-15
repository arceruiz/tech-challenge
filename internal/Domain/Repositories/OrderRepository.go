package repositories

import (
	"database/sql"
	entities "tech-challenge/internal/Domain/Entities"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) GetOrders() ([]entities.Order, error) {
	return nil, nil
}

func (r *OrderRepository) CreateOrder(order entities.Order) (entities.Order, error) {
	return entities.Order{}, nil
}

func (r *OrderRepository) GetOrder(id string) (entities.Order, error) {
	return entities.Order{}, nil
}

func (r *OrderRepository) UpdateOrder(id string, updatedOrder entities.Order) (entities.Order, error) {
	return entities.Order{}, nil
}

func (r *OrderRepository) DeleteOrder(id string) error {
	return nil
}
