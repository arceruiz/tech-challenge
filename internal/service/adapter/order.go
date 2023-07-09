package adapter

import (
	"tech-challenge/internal/canonical"
	"tech-challenge/internal/repository/adapters"
	repos "tech-challenge/internal/repository/port"
	services "tech-challenge/internal/service/port"

	"github.com/google/uuid"
)

type orderService struct {
	repo repos.OrderRepository
}

func NewOrderService() services.OrderService {
	return &orderService{
		adapters.NewOrderRepo(),
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
