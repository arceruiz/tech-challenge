package usecases

import (
	entities "tech-challenge/internal/Domain/Entities"
	repositories "tech-challenge/internal/Domain/Repositories"
	"time"
)

type CustomerUseCase struct {
	repo *repositories.CustomerRepository
}

func NewCustomerUseCase(repo *repositories.CustomerRepository) *CustomerUseCase {
	return &CustomerUseCase{
		repo: repo,
	}
}

func (s *CustomerUseCase) GetCustomers() ([]entities.Customer, error) {
	return s.repo.GetCustomers()
}

func (s *CustomerUseCase) CreateCustomer(customer entities.Customer) (entities.Customer, error) {
	customer.CreatedAt = time.Now()
	return s.repo.CreateCustomer(customer)
}

func (s *CustomerUseCase) GetCustomer(id string) (entities.Customer, error) {
	return s.repo.GetCustomer(id)
}

func (s *CustomerUseCase) UpdateCustomer(id string, updatedCustomer entities.Customer) (entities.Customer, error) {
	return s.repo.UpdateCustomer(id, updatedCustomer)
}

func (s *CustomerUseCase) DeleteCustomer(id string) error {
	return s.repo.DeleteCustomer(id)
}
