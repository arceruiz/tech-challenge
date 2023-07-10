package adapters

import (
	"database/sql"
	"strconv"
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
	orderRows, err := r.db.Query(
		"SELECT * FROM Order",
	)
	if err != nil {
		return nil, err
	}
	defer orderRows.Close()

	var orders []canonical.Order

	for orderRows.Next() {
		var order canonical.Order
		if err = orderRows.Scan(
			&order.ID,
			&order.CustomerID,
			&order.PaymentID,
			&order.Status,
			&order.CreatedAt,
			&order.UpdatedAt,
			&order.Total,
		); err != nil {
			return nil, err
		}
		orderItems, err := r.GetOrderItemsFromOrderID(order.ID)
		if err != nil {
			return nil, err
		}

		productRepo := NewProductRepo()
		var cannonOrderItems []canonical.OrderItem
		for _, orderItem := range orderItems {
			product, err := productRepo.GetByID(orderItem.ProductID)
			if err != nil {
				return nil, err
			}

			qty, err := strconv.Atoi(orderItem.Quantity)
			if err != nil {
				return nil, err
			}

			cannonOrderItems = append(cannonOrderItems, canonical.OrderItem{
				Product:  *product,
				Quantity: qty,
			})
		}

		order.OrderItems = cannonOrderItems

		orders = append(orders, order)
	}

	return orders, nil
}

func (r *orderRepository) CreateOrder(order canonical.Order) error {
	sqlStatement := "INSERT INTO Order (ID, CustomerID, PaymentID, Status , CreatedAt , UpdatedAt , Total) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	_, err := r.db.Exec(sqlStatement, order.ID, order.CustomerID, order.PaymentID, order.Status, order.CreatedAt, order.UpdatedAt, order.Total)
	if err != nil {
		return err
	}
	for _, orderItem := range order.OrderItems {
		err := r.CreateOrderItem(orderItem, order.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *orderRepository) UpdateOrder(id string, updatedOrder canonical.Order) (canonical.Order, error) {
	return canonical.Order{}, nil
}

func (r *orderRepository) GetByID(id string) (*canonical.Order, error) {
	orderRows, err := r.db.Query(
		"SELECT * FROM Order WHERE ID = ?",
		id,
	)
	if err != nil {
		return nil, err
	}
	defer orderRows.Close()

	var order *canonical.Order
	if orderRows.Next() {
		if err = orderRows.Scan(
			&order.ID,
			&order.CustomerID,
			&order.PaymentID,
			&order.Status,
			&order.CreatedAt,
			&order.UpdatedAt,
			&order.Total,
		); err != nil {
			return nil, err
		}
		orderItems, err := r.GetOrderItemsFromOrderID(order.ID)
		if err != nil {
			return nil, err
		}

		productRepo := NewProductRepo()
		var cannonOrderItems []canonical.OrderItem
		for _, orderItem := range orderItems {
			product, err := productRepo.GetByID(orderItem.ProductID)
			if err != nil {
				return nil, err
			}

			qty, err := strconv.Atoi(orderItem.Quantity)
			if err != nil {
				return nil, err
			}

			cannonOrderItems = append(cannonOrderItems, canonical.OrderItem{
				Product:  *product,
				Quantity: qty,
			})
		}

		order.OrderItems = cannonOrderItems

	}

	return order, nil
}

func (r *orderRepository) GetByStatus(status string) ([]canonical.Order, error) {
	orderRows, err := r.db.Query(
		"SELECT * FROM Order WHERE Status = ?",
		status,
	)
	if err != nil {
		return nil, err
	}
	defer orderRows.Close()

	var orders []canonical.Order

	for orderRows.Next() {
		var order canonical.Order
		if err = orderRows.Scan(
			&order.ID,
			&order.CustomerID,
			&order.PaymentID,
			&order.Status,
			&order.CreatedAt,
			&order.UpdatedAt,
			&order.Total,
		); err != nil {
			return nil, err
		}
		orderItems, err := r.GetOrderItemsFromOrderID(order.ID)
		if err != nil {
			return nil, err
		}

		productRepo := NewProductRepo()
		var cannonOrderItems []canonical.OrderItem
		for _, orderItem := range orderItems {
			product, err := productRepo.GetByID(orderItem.ProductID)
			if err != nil {
				return nil, err
			}

			qty, err := strconv.Atoi(orderItem.Quantity)
			if err != nil {
				return nil, err
			}

			cannonOrderItems = append(cannonOrderItems, canonical.OrderItem{
				Product:  *product,
				Quantity: qty,
			})
		}

		order.OrderItems = cannonOrderItems

		orders = append(orders, order)
	}

	return orders, nil
}

func (r *orderRepository) CheckoutOrder(orderID string, payment canonical.Payment) error {
	sqlStatement := "INSERT INTO Order (ID, PaymentType, CreatedAt) VALUES ($1, $2, $3)"

	_, err := r.db.Exec(sqlStatement, payment.ID, payment.PaymentType, payment.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
