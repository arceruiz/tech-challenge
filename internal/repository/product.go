package repository

import (
	"context"
	"tech-challenge/internal/canonical"

	"github.com/jackc/pgx/v4/pgxpool"
)

type ProductRepository interface {
	GetProducts(context.Context) ([]canonical.Product, error)
	CreateProduct(context.Context, canonical.Product) (int, error)
	UpdateProduct(context.Context, string, canonical.Product) error
	GetByID(context.Context, string) (*canonical.Product, error)
	GetByCategory(context.Context, string) ([]canonical.Product, error)
}

type productRepository struct {
	db *pgxpool.Pool
}

func NewProductRepo() ProductRepository {
	return &productRepository{New()}
}

func (r *productRepository) GetProducts(ctx context.Context) ([]canonical.Product, error) {
	rows, err := r.db.Query(ctx,
		"SELECT * FROM \"Product\" WHERE Status = 2",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []canonical.Product

	for rows.Next() {
		var product canonical.Product
		if err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Category,
			&product.Status,
			&product.ImagePath,
		); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *productRepository) CreateProduct(ctx context.Context, product canonical.Product) (int, error) {
	sqlStatement := "INSERT INTO \"Product\" (Name, Description, Price, Category, Status, Imagepath) VALUES ($1, $2, $3, $4, $5, $6) RETURNING ID"
	var insertedId int

	err := r.db.QueryRow(ctx, sqlStatement, product.Name, product.Description, product.Price, product.Category, product.Status, product.ImagePath).Scan(&insertedId)
	if err != nil {
		return 0, err
	}
	return insertedId, nil
}

func (r *productRepository) UpdateProduct(ctx context.Context, id string, product canonical.Product) error {
	sqlStatement := "UPDATE \"Product\" SET Name = $1, Description = $2, Price = $3, Category = $4, Status = $5, Imagepath = $6 WHERE ID = $7"

	_, err := r.db.Exec(ctx, sqlStatement, product.Name, product.Description, product.Price, product.Category, product.Status, product.ImagePath, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepository) GetByID(ctx context.Context, id string) (*canonical.Product, error) {
	rows, err := r.db.Query(ctx,
		"SELECT * FROM \"Product\" WHERE ID = $1",
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var product canonical.Product
	if rows.Next() {
		if err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Category,
			&product.Status,
			&product.ImagePath,
		); err != nil {
			return nil, err
		}
		return &product, nil
	}

	return nil, ErrorNotFound
}

func (r *productRepository) GetByCategory(ctx context.Context, Category string) ([]canonical.Product, error) {
	rows, err := r.db.Query(ctx,
		"SELECT * FROM \"Product\" WHERE Category = $1",
		Category,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []canonical.Product

	for rows.Next() {
		var product canonical.Product
		if err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Category,
			&product.Status,
			&product.ImagePath,
		); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
