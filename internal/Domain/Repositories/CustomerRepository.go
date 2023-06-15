package repositories

import (
	"database/sql"
	entities "tech-challenge/internal/Domain/Entities"
)

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func (r *CustomerRepository) GetCustomers() ([]entities.Customer, error) {
	return nil, nil
}

func (r *CustomerRepository) CreateCustomer(customer entities.Customer) (entities.Customer, error) {
	return entities.Customer{}, nil
}

func (r *CustomerRepository) GetCustomer(id string) (entities.Customer, error) {
	return entities.Customer{}, nil
}

func (r *CustomerRepository) UpdateCustomer(id string, updatedCustomer entities.Customer) (entities.Customer, error) {
	return entities.Customer{}, nil
}

func (r *CustomerRepository) DeleteCustomer(id string) error {
	return nil
}
