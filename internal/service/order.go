package service

import (
	"context"
	"tech-challenge/internal/canonical"
	"tech-challenge/internal/repository"
	"time"

	"github.com/shopspring/decimal"
)

type OrderService interface {
	GetOrders(context.Context) ([]canonical.Order, error)
	CreateOrder(context.Context, canonical.Order) (int, error)
	UpdateOrder(context.Context, string, canonical.Order) (canonical.Order, error)
	GetByID(context.Context, string) (*canonical.Order, error)
	GetByStatus(context.Context, string) ([]canonical.Order, error)
	CheckoutOrder(context.Context, string, canonical.Payment) error
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService() OrderService {
	return &orderService{
		repository.NewOrderRepo(),
	}
}

func (s *orderService) GetOrders(ctx context.Context) ([]canonical.Order, error) {
	return s.repo.GetOrders(ctx)
}

func (s *orderService) CreateOrder(ctx context.Context, order canonical.Order) (int, error) {
	timeNow := time.Now()
	order.CreatedAt = &timeNow
	order.Status = canonical.ORDER_RECEIVED
	s.calculateTotal(&order)

	return s.repo.CreateOrder(ctx, order)
}

func (s *orderService) UpdateOrder(ctx context.Context, id string, updatedOrder canonical.Order) (canonical.Order, error) {
	return s.repo.UpdateOrder(ctx, id, updatedOrder)
}

func (s *orderService) GetByID(ctx context.Context, id string) (*canonical.Order, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *orderService) GetByStatus(ctx context.Context, id string) ([]canonical.Order, error) {
	return s.repo.GetByStatus(ctx, id)
}

func (s *orderService) CheckoutOrder(ctx context.Context, orderID string, payment canonical.Payment) error {
	order.Status = canonical.ORDER_PREPARING
	order.UpdatedAt = time.Now()
	err := s.repo.CheckoutOrder(ctx, orderID, payment)
	if err != nil {
		return err
	}

	return nil
}

func (s *orderService) calculateTotal(order *canonical.Order) {
	for _, product := range order.OrderItems {
		price := decimal.NewFromFloat(product.Price)
		quantity := decimal.NewFromInt(product.Quantity)
		productTotal, _ := price.Mul(quantity).Float64()

		order.Total += productTotal
	}
}
