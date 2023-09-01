package service

import (
	"context"
	"tech-challenge/internal/auth/token"
	"tech-challenge/internal/canonical"
	"tech-challenge/internal/repository"
	"time"

	"fmt"
	"tech-challenge/internal/security"

	"github.com/sirupsen/logrus"
)

type CustomerService interface {
	Create(context.Context, canonical.Customer) (*canonical.Customer, error)
	Login(context.Context, canonical.Customer) (string, error)
	Bypass() (string, error)
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

func (u *customerService) Login(ctx context.Context, customer canonical.Customer) (string, error) {
	customerBase, err := u.repo.GetByEmail(ctx, customer.Email)
	if err != nil {
		err = fmt.Errorf("error getting customer by email: %w", err)
		logrus.WithError(err).Info()
		return "", err
	}

	if err = security.CheckPassword(customerBase.Password, customer.Password); err != nil {
		err = fmt.Errorf("error checking password: %w", err)
		logrus.WithError(err).Info()
		return "", err
	}

	token, err := token.GenerateToken(customerBase.Id)
	if err != nil {
		err = fmt.Errorf("error generting token: %w", err)
		logrus.WithField("customerId", customerBase.Id).WithError(err).Warn()
		return "", err
	}

	return token, nil
}

func (u *customerService) Bypass() (string, error) {
	token, err := token.GenerateToken(1)
	if err != nil {
		err = fmt.Errorf("error generting token: %w", err)
		logrus.WithError(err).Warn()
		return "", err
	}

	return token, nil
}
