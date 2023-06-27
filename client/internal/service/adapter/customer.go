package adapter

import (
	"client/internal/auth/token"
	"client/internal/canonical"
	"client/internal/repository/adapters"
	"client/internal/repository/ports"
	"client/internal/security"
	"fmt"

	"github.com/google/uuid"
)

type CustomerService interface {
	Create(canonical.Customer) (canonical.Customer, error)
	Login(user canonical.Customer) (string, error)
	Bypass()
}

type customerService struct {
	repo ports.CustomerRepository
}

func NewCustomerService() CustomerService {
	return &customerService{
		adapters.NewCustomerRepo(),
	}
}

func (u *customerService) Create(user canonical.Customer) (canonical.Customer, error) {
	passEncrypted, err := security.Hash(user.Password)
	if err != nil {
		return canonical.Customer{}, err
	}
	user.Password = string(passEncrypted)
	user.Id = uuid.NewString()

	err = u.repo.Create(user)
	if err != nil {
		return canonical.Customer{}, err
	}

	return user, nil
}

func (u *customerService) Login(user canonical.Customer) (string, error) {
	userBase, err := u.repo.GetByEmail(user.Email)
	if err != nil {
		return "", fmt.Errorf("error getting customer by email: %w", err)
	}

	if err = security.CheckPassword(userBase.Password, user.Password); err != nil {
		return "", err
	}

	token, err := token.GenerateToken(userBase.Id)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *customerService) Bypass() {

}
