package adapter

import (
	"tech-challenge/internal/canonical"
	"tech-challenge/internal/repository/adapters"
	"tech-challenge/internal/repository/port"
)

type OrderService interface {
	GetOrders() ([]canonical.Order, error)
	CreateOrder(order canonical.Order) (canonical.Order, error)
	UpdateOrder(id string, updatedOrder canonical.Order) (canonical.Order, error)
	GetByID(id string) (canonical.Order, error)
	GetByStatus(id string) ([]canonical.Order, error)
}

type orderService struct {
	repo port.OrderRepository
}

func NewOrderService() OrderService {
	return &orderService{
		adapters.NewOrderRepo(),
	}
}

func (s *orderService) GetOrders() ([]canonical.Order, error) {
	return s.repo.GetOrders()
}

func (s *orderService) CreateOrder(order canonical.Order) (canonical.Order, error) {
	return s.repo.CreateOrder(order)
}

func (s *orderService) UpdateOrder(id string, updatedOrder canonical.Order) (canonical.Order, error) {
	return s.repo.UpdateOrder(id, updatedOrder)
}

func (s *orderService) GetByID(id string) (canonical.Order, error) {
	return s.repo.GetByID(id)
}

func (s *orderService) GetByStatus(id string) ([]canonical.Order, error) {
	return s.repo.GetByStatus(id)
}
