package service

import (
	"context"
	"tech-challenge/internal/canonical"
	"tech-challenge/internal/repository"
	"time"

	"fmt"
	"tech-challenge/internal/security"

	"github.com/sirupsen/logrus"
)

type CustomerService interface {
	Create(context.Context, canonical.Customer) (*canonical.Customer, error)
	Get(context.Context, canonical.Customer) (*canonical.Customer, error)
}

type customerService struct {
	repo repository.CustomerRepository
}

func NewCustomerService() CustomerService {
	return &customerService{
		repository.NewCustomerRepo(),
	}
}

func (u *customerService) Create(ctx context.Context, customer canonical.Customer) (*canonical.Customer, error) {
	createdAt := time.Now()
	customer.CreatedAt = &createdAt

	passEncrypted, err := security.Hash(customer.Password)
	if err != nil {
		err = fmt.Errorf("error generating password hash: %w", err)
		logrus.WithError(err).Warn()
		return nil, err
	}
	customer.Password = string(passEncrypted)

	id, err := u.repo.Create(ctx, customer)
	if err != nil {
		err = fmt.Errorf("error saving customer in database: %w", err)
		logrus.WithError(err).Warn()
		return nil, err
	}

	customer.Id = id
	return &customer, nil
}

func (u *customerService) Get(ctx context.Context, customer canonical.Customer) (*canonical.Customer, error) {
	if customer.Document != "" {
		baseCustomer, err := u.repo.GetByDocument(ctx, customer.Document)
		if err != nil {
			return nil, err
		}

		return baseCustomer, nil
	}

	baseCustomer, err := u.repo.GetByEmail(ctx, customer.Email)
	if err != nil {
		return nil, err
	}

	return baseCustomer, nil
}
