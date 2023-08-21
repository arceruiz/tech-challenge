package repository

import (
	"database/sql"
	"tech-challenge/internal/canonical"
)

type ProductRepository interface {
	GetProducts() ([]canonical.Product, error)
	CreateProduct(product canonical.Product) error
	UpdateProduct(id string, product canonical.Product) error
	GetByID(id string) (*canonical.Product, error)
	GetByCategory(id string) ([]canonical.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepo() ProductRepository {
	return &productRepository{New()}
}

func (r *productRepository) GetProducts() ([]canonical.Product, error) {
	rows, err := r.db.Query(
		"SELECT * FROM Product WHERE Status = 'ACTIVE'",
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

func (r *productRepository) CreateProduct(product canonical.Product) error {
	sqlStatement := "INSERT INTO Product (ID, Name, Description, Price, Category, Status, Imagepath) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	_, err := r.db.Exec(sqlStatement, product.ID, product.Name, product.Description, product.Price, product.Category, product.Status, product.ImagePath)
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepository) UpdateProduct(id string, product canonical.Product) error {
	sqlStatement := "UPDATE Product SET Name = ?, Description = ?, Price = ?, Category = ?, Status = ?, Imagepath = ? WHERE ID = ?"

	_, err := r.db.Exec(sqlStatement, product.Name, product.Description, product.Price, product.Category, product.Status, product.ImagePath, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepository) GetByID(id string) (*canonical.Product, error) {
	rows, err := r.db.Query(
		"SELECT * FROM Product WHERE ID = $1",
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

func (r *productRepository) GetByCategory(Category string) ([]canonical.Product, error) {
	rows, err := r.db.Query(
		"SELECT * FROM Product WHERE Category = $1",
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
