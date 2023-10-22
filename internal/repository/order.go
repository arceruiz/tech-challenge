package repository

import (
	"context"
	"tech-challenge/internal/canonical"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type OrderRepository interface {
	GetOrders(context.Context) ([]canonical.Order, error)
	CreateOrder(context.Context, canonical.Order) (int, error)
	UpdateOrder(context.Context, string, canonical.Order) error
	GetByID(context.Context, string) (*canonical.Order, error)
	GetByStatus(context.Context, string) ([]canonical.Order, error)
	CheckoutOrder(context.Context, string, canonical.Payment) (int, error)
}

type orderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepo() OrderRepository {
	return &orderRepository{New()}
}

func (r *orderRepository) GetOrders(ctx context.Context) ([]canonical.Order, error) {
	sqlStatement :=
		`
		SELECT 
			* 
		FROM 
			"Order" 
		WHERE
			Status < 3
		ORDER BY
			Status, 
			createdat
		`

	orderRows, err := r.db.Query(ctx,
		sqlStatement,
	)
	if err != nil {
		return nil, err
	}
	defer orderRows.Close()

	var orders []canonical.Order

	for orderRows.Next() {
		var orderDTO Order
		err = orderRows.Scan(
			&orderDTO.ID,
			&orderDTO.CustomerID,
			&orderDTO.PaymentID,
			&orderDTO.Status,
			&orderDTO.CreatedAt,
			&orderDTO.UpdatedAt,
			&orderDTO.Total,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, orderDTO.toCanonical())
	}

	return orders, nil
}

func (r *orderRepository) CreateOrder(ctx context.Context, order canonical.Order) (int, error) {
	sqlStatementOrder := "INSERT INTO \"Order\" (CustomerID, Status, CreatedAt, UpdatedAt, Total) VALUES ($1, $2, $3, $4, $5) RETURNING ID"

	var insertedId int

	err := r.db.QueryRow(ctx, sqlStatementOrder, order.CustomerID, order.Status, order.CreatedAt, order.UpdatedAt, order.Total).Scan(&insertedId)
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
	sqlStatement := "UPDATE \"Order\" SET CustomerID = $1, PaymentID = $2, Status = $3, CreatedAt = $4, UpdatedAt = $5, Total = $6 WHERE ID = $7"

	_, err := r.db.Exec(ctx, sqlStatement, updatedOrder.CustomerID, updatedOrder.Payment.ID, updatedOrder.Status, updatedOrder.CreatedAt, updatedOrder.UpdatedAt, updatedOrder.Total, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *orderRepository) GetByID(ctx context.Context, id string) (*canonical.Order, error) {
	var orderDTO OrderJoinPayment
	var order *canonical.Order
	var batch pgx.Batch

	batch.Queue(`SELECT o.*, p.PaymentType, p.createdat AS paymentCreatedat, p.status as paymentstatus FROM "Order" o LEFT JOIN "Payment" p ON p.Id = o.paymentid WHERE o.ID = $1;`, id)
	batch.Queue(`SELECT p.id, p.name, p.description, p.price, p.category, p.status, p.imagepath, oi.quantity FROM "Order_Items" oi JOIN "Product" p ON p.ID = oi.productid WHERE oi.orderid = $1;`, id)

	results := r.db.SendBatch(ctx, &batch)

	orderRow, err := results.Query()
	defer orderRow.Close()
	if err != nil {
		return nil, err
	}

	if orderRow.Next() {
		if err = orderRow.Scan(
			&orderDTO.ID,
			&orderDTO.CustomerID,
			&orderDTO.PaymentID,
			&orderDTO.Status,
			&orderDTO.CreatedAt,
			&orderDTO.UpdatedAt,
			&orderDTO.Total,
			&orderDTO.PaymentType,
			&orderDTO.PaymentCreatedat,
			&orderDTO.PaymentStatus,
		); err != nil {
			return nil, err
		}

		dt := orderDTO.toCanonical()

		order = &dt
	}
	if order == nil {
		return nil, nil
	}

	orderRow.Close()

	productsRows, err := results.Query()
	defer productsRows.Close()
	if err != nil {
		return nil, err
	}
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
	productsRows.Close()

	return order, nil
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
			&order.Payment.ID,
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

func (r *orderRepository) CheckoutOrder(ctx context.Context, orderID string, payment canonical.Payment) (int, error) {
	sqlStatement := "INSERT INTO \"Payment\" (PaymentType, CreatedAt, Status) VALUES ($1, $2, $3) RETURNING ID"

	var insertedId int
	err := r.db.QueryRow(ctx, sqlStatement, payment.PaymentType, payment.CreatedAt, payment.Status).Scan(&insertedId)
	if err != nil {
		return 0, err
	}
	return insertedId, nil
}
