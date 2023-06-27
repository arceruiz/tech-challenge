package repository

import (
	"client/internal/canonical"
	"database/sql"

	"github.com/sirupsen/logrus"
)

type OrderRepository interface {
	GetOrders() ([]canonical.Order, error)
	CreateOrder(order canonical.Order) (canonical.Order, error)
	UpdateOrder(id string, updatedOrder canonical.Order) (canonical.Order, error)
	GetByID(id string) (canonical.Order, error)
	GetByStatus(id string) ([]canonical.Order, error)
}

type orderRepository struct {
	db *sql.DB
}

func NewOrderRepo() OrderRepository {
	connStr := "host=localhost port=5432 dbname=fiap_tech_challenge user=postgres password=1234 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logrus.Fatal(err)
	}
	return &orderRepository{db}
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
