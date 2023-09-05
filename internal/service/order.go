package service

import (
	"context"
	"fmt"
	"strconv"
	"tech-challenge/internal/canonical"
	"tech-challenge/internal/repository"
	"time"

	"github.com/shopspring/decimal"
)

type OrderService interface {
	GetOrders(context.Context) ([]canonical.Order, error)
	CreateOrder(context.Context, canonical.Order) (int, error)
	UpdateOrder(context.Context, string, canonical.Order) error
	GetByID(context.Context, string) (*canonical.Order, error)
	GetByStatus(context.Context, string) ([]canonical.Order, error)
	CheckoutOrder(context.Context, string, canonical.Payment) (*canonical.Order, error)
	PaymentCallback(ctx context.Context, orderID, status string) error
}

type orderService struct {
	repo           repository.OrderRepository
	paymentService PaymentService
}

func NewOrderService() OrderService {
	return &orderService{
		repo:           repository.NewOrderRepo(),
		paymentService: NewPaymentService(),
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

func (s *orderService) UpdateOrder(ctx context.Context, id string, updatedOrder canonical.Order) error {
	return s.repo.UpdateOrder(ctx, id, updatedOrder)
}

func (s *orderService) GetByID(ctx context.Context, id string) (*canonical.Order, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *orderService) GetByStatus(ctx context.Context, id string) ([]canonical.Order, error) {
	return s.repo.GetByStatus(ctx, id)
}

func (s *orderService) CheckoutOrder(ctx context.Context, orderID string, payment canonical.Payment) (*canonical.Order, error) {
	order, err := s.repo.GetByID(ctx, orderID)
	if err != nil {
		return nil, fmt.Errorf("payment not criated, error searching order, %w", err)
	}
	now := time.Now()

	payment.CreatedAt = &now
	payment.Status = canonical.PAYMENT_INIT
	paymentId, err := s.repo.CheckoutOrder(ctx, orderID, payment)
	if err != nil {
		return nil, fmt.Errorf("error checking out order, %w", err)
	}
	order.Payment = &canonical.Payment{
		ID: paymentId,
	}
	order.Status = canonical.ORDER_PREPARING
	order.UpdatedAt = &now
	err = s.repo.UpdateOrder(ctx, orderID, *order)
	if err != nil {
		return nil, fmt.Errorf("payment not criated, error updating order, %w", err)
	}

	order, err = s.repo.GetByID(ctx, orderID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving updated order, %w", err)
	}

	return order, nil
}

func (s *orderService) PaymentCallback(ctx context.Context, orderID, status string) error {

	order, err := s.GetByID(ctx, orderID)
	if err != nil {
		return fmt.Errorf("payment not updated, error searching order, %w", err)
	}

	order.Status = s.mapPaymentCallbackStatus(status)
	now := time.Now()
	order.UpdatedAt = &now

	s.UpdateOrder(ctx, orderID, *order)
	if err != nil {
		return fmt.Errorf("payment not updated, error updating order, %w", err)
	}

	s.paymentService.Callback(ctx, strconv.Itoa(order.Payment.ID), status)

	return nil
}

func (s *orderService) mapPaymentCallbackStatus(status string) canonical.OrderStatus {
	switch status {
	case "ERROR":
		return canonical.ORDER_CANCELLED
	case "COMPLETED":
		return canonical.ORDER_PREPARING
	case "PENDING":
		return canonical.ORDER_RECEIVED
	default:
		return canonical.ORDER_RECEIVED
	}
}

func (s *orderService) calculateTotal(order *canonical.Order) {
	for _, product := range order.OrderItems {
		price := decimal.NewFromFloat(product.Price)
		quantity := decimal.NewFromInt(product.Quantity)
		productTotal, _ := price.Mul(quantity).Float64()

		order.Total += productTotal
	}
}
