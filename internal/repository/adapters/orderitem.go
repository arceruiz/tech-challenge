package adapters

import (
	"strconv"
	"tech-challenge/internal/canonical"
	"tech-challenge/internal/repository"

	"github.com/google/uuid"
)

func (r *orderRepository) GetOrderItemsFromOrderID(orderID string) ([]repository.OrderItem, error) {
	orderItemRows, err := r.db.Query(
		"SELECT * FROM OrderItem WEHERE OrderID = ?",
		orderID,
	)
	if err != nil {
		return nil, err
	}
	defer orderItemRows.Close()

	var orderItems []repository.OrderItem

	for orderItemRows.Next() {
		var orderItem repository.OrderItem
		if err = orderItemRows.Scan(
			&orderItem.ID,
			&orderItem.OrderID,
			&orderItem.ProductID,
			&orderItem.Quantity,
		); err != nil {
			return nil, err
		}
		orderItems = append(orderItems, orderItem)
	}

	return orderItems, nil
}

func (r *orderRepository) CreateOrderItem(cannonOrderItem canonical.OrderItem, orderID string) error {
	sqlStatement := "INSERT INTO Product (ID, OrderID, ProductID, Quantity) VALUES ($1, $2, $3, $4)"
	orderItem := mapCannonOrderItem(cannonOrderItem, orderID)
	_, err := r.db.Exec(sqlStatement, orderItem.ID, orderItem.OrderID, orderItem.ProductID, orderItem.Quantity)
	if err != nil {
		return err
	}
	return nil
}

func mapCannonOrderItem(cannonOrderItem canonical.OrderItem, orderID string) repository.OrderItem {
	return repository.OrderItem{
		ID:        uuid.NewString(),
		OrderID:   orderID,
		ProductID: cannonOrderItem.Product.ID,
		Quantity:  strconv.Itoa(cannonOrderItem.Quantity),
	}
}
