package service

import (
	"tech-challenge/internal/canonical"
	"tech-challenge/internal/repository"

	"github.com/google/uuid"
)

type OrderService interface {
	GetOrders() ([]canonical.Order, error)
	CreateOrder(order canonical.Order) error
	UpdateOrder(id string, updatedOrder canonical.Order) (canonical.Order, error)
	GetByID(id string) (*canonical.Order, error)
	GetByStatus(id string) ([]canonical.Order, error)
	CheckoutOrder(orderID string, payment canonical.Payment) error
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService() OrderService {
	return &orderService{
		repository.NewOrderRepo(),
	}
}

func (s *orderService) GetOrders() ([]canonical.Order, error) {
	return s.repo.GetOrders()
}

func (s *orderService) CreateOrder(order canonical.Order) error {
	order.ID = uuid.NewString()
	return s.repo.CreateOrder(order)
}

func (s *orderService) UpdateOrder(id string, updatedOrder canonical.Order) (canonical.Order, error) {
	return s.repo.UpdateOrder(id, updatedOrder)
}

func (s *orderService) GetByID(id string) (*canonical.Order, error) {
	return s.repo.GetByID(id)
}

func (s *orderService) GetByStatus(id string) ([]canonical.Order, error) {
	return s.repo.GetByStatus(id)
}

func (s *orderService) CheckoutOrder(orderID string, payment canonical.Payment) error {
	err := s.repo.CheckoutOrder(orderID, payment)
	if err != nil {
		return err
	}

	return nil
}
