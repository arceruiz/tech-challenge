package repository

import (
	"context"
	"tech-challenge/internal/canonical"

	"github.com/jackc/pgx/v4"
)

type OrderRepository interface {
	GetOrders(context.Context) ([]canonical.Order, error)
	CreateOrder(context.Context, canonical.Order) (int, error)
	UpdateOrder(context.Context, string, canonical.Order) error
	GetByID(context.Context, string) (*canonical.Order, error)
	GetByStatus(context.Context, string) ([]canonical.Order, error)
	CheckoutOrder(context.Context, string, canonical.Payment) error
}

type orderRepository struct {
	db *pgx.Conn
}

func NewOrderRepo() OrderRepository {
	return &orderRepository{New()}
}

func (r *orderRepository) GetOrders(ctx context.Context) ([]canonical.Order, error) {
	sqlStatement := "SELECT * FROM \"Order\" where Status != 4 ORDER BY Status DESC"

	orderRows, err := r.db.Query(ctx,
		sqlStatement,
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
		orders = append(orders, order)
	}

	return orders, nil
}

func (r *orderRepository) CreateOrder(ctx context.Context, order canonical.Order) (int, error) {
	sqlStatementOrder := "INSERT INTO \"Order\" (CustomerID, PaymentID, Status, CreatedAt, UpdatedAt, Total) VALUES ($1, $2, $3, $4, $5, $6) RETURNING ID"

	var insertedId int

	err := r.db.QueryRow(ctx, sqlStatementOrder, order.CustomerID, order.PaymentID, order.Status, order.CreatedAt, order.UpdatedAt, order.Total).Scan(&insertedId)
	if err != nil {
		return 0, err
	}

	if len(order.OrderItems) > 0 {
		sqlStatementOrderProduct := "INSERT INTO \"Order_Items\" (OrderID, ProductId, Quantity) VALUES ($1, $2, $3)"

		for _, product := range order.OrderItems {
			_, err = r.db.Exec(ctx, sqlStatementOrderProduct, insertedId, product.ID, product.Quantity)
			if err != nil {
				return 0, err
			}
		}
	}

	return insertedId, nil
}

func (r *orderRepository) UpdateOrder(ctx context.Context, id string, updatedOrder canonical.Order) error {
	sqlStatement := "UPDATE \"Order\" SET CustomerID = ?, PaymentID = ?, Status = ?, CreatedAt = ?, UpdatedAt = ?, Total = ? WHERE ID = ?"

	_, err := r.db.Exec(ctx, sqlStatement, updatedOrder.CustomerID, updatedOrder.PaymentID, updatedOrder.Status, updatedOrder.CreatedAt, updatedOrder.UpdatedAt, updatedOrder.Total, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *orderRepository) GetByID(ctx context.Context, id string) (*canonical.Order, error) {
	var order canonical.Order
	var batch pgx.Batch

	batch.Queue(`SELECT * FROM "Order" o WHERE o.ID = $1;`, id)
	batch.Queue(`SELECT p.id, p.name, p.description, p.price, p.category, p.status, p.imagepath, oi.quantity FROM "Order_Items" oi JOIN "Product" p ON p.ID = oi.productid WHERE oi.orderid = $1;`, id)

	results := r.db.SendBatch(ctx, &batch)

	orderRow, err := results.Query()
	if err != nil {
		return nil, err
	}

	if orderRow.Next() {
		err = orderRow.Scan(
			&order.ID,
			&order.CustomerID,
			&order.PaymentID,
			&order.Status,
			&order.CreatedAt,
			&order.UpdatedAt,
			&order.Total,
		)
		if err != nil {
			return nil, err
		}
	}
	orderRow.Close()

	productsRows, err := results.Query()
	if err != nil {
		return nil, err
	}
	defer productsRows.Close()

	var items []canonical.OrderItem

	for productsRows.Next() {
		var item canonical.OrderItem

		err = productsRows.Scan(
			&item.ID,
			&item.Name,
			&item.Description,
			&item.Price,
			&item.Category,
			&item.Status,
			&item.ImagePath,
			&item.Quantity,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	order.OrderItems = items
	return &order, nil
}

func (r *orderRepository) GetByStatus(ctx context.Context, status string) ([]canonical.Order, error) {
	sqlStatement := "SELECT * FROM \"Order\" WHERE status = $1"

	orderRows, err := r.db.Query(ctx,
		sqlStatement,
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
		orders = append(orders, order)
	}

	return orders, nil
}

func (r *orderRepository) CheckoutOrder(ctx context.Context, orderID string, payment canonical.Payment) error {
	sqlStatement := "INSERT INTO \"Payment\" (ID, PaymentType, CreatedAt) VALUES ($1, $2, $3)"

	_, err := r.db.Exec(ctx, sqlStatement, payment.ID, payment.PaymentType, payment.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
