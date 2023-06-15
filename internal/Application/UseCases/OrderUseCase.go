package usecases

import (
	entities "tech-challenge/internal/Domain/Entities"
	repositories "tech-challenge/internal/Domain/Repositories"
	"time"
)

type OrderUseCase struct {
	repo *repositories.OrderRepository
}

func NewOrderUseCase(repo *repositories.OrderRepository) *OrderUseCase {
	return &OrderUseCase{
		repo: repo,
	}
}

func (s *OrderUseCase) GetOrders() ([]entities.Order, error) {
	return s.repo.GetOrders()
}

func (s *OrderUseCase) CreateOrder(order entities.Order) (entities.Order, error) {
	order.CreatedAt = time.Now()
	return s.repo.CreateOrder(order)
}

func (s *OrderUseCase) GetOrder(id string) (entities.Order, error) {
	return s.repo.GetOrder(id)
}

func (s *OrderUseCase) UpdateOrder(id string, updatedOrder entities.Order) (entities.Order, error) {
	return s.repo.UpdateOrder(id, updatedOrder)
}

func (s *OrderUseCase) DeleteOrder(id string) error {
	return s.repo.DeleteOrder(id)
}
