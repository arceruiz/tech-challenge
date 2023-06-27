package service

import (
	"client/internal/auth/token"
	"client/internal/canonical"
	"client/internal/repository"
	"client/internal/security"
	"fmt"

	"github.com/google/uuid"
)

type UserService interface {
	Register(canonical.User) (canonical.User, error)
	Login(user canonical.User) (string, error)
	Bypass()
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService() UserService {
	return &userService{
		repository.NewUserRepo(),
	}
}

func (u *userService) Register(user canonical.User) (canonical.User, error) {
	passEncrypted, err := security.Hash(user.Password)
	if err != nil {
		return canonical.User{}, err
	}
	user.Password = string(passEncrypted)
	user.Id = uuid.NewString()

	err = u.repo.Create(user)
	if err != nil {
		return canonical.User{}, err
	}

	return user, nil
}

func (u *userService) Login(user canonical.User) (string, error) {
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

func (u *userService) Bypass() {

}
