package rest

import (
	"fmt"
	"tech-challenge/internal/service"

	"net/http"

	"github.com/labstack/echo/v4"
)

type Customer interface {
	RegisterGroup(*echo.Group)
	Create(echo.Context) error
	Login(echo.Context) error
}

type customer struct {
	service service.CustomerService
}

func NewCustomerChannel() Customer {
	return &customer{
		service: service.NewCustomerService(),
	}
}

func (u *customer) RegisterGroup(g *echo.Group) {
	g.POST("/create", u.Create)
	g.POST("/login", u.Login)
	g.POST("/bypass", u.Bypass)
}

func (u *customer) Create(c echo.Context) error {
	var customerRequest CustomerRequest

	if err := c.Bind(&customerRequest); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Message: fmt.Errorf("invalid data").Error(),
		})
	}

	customer, err := u.service.Create(c.Request().Context(), customerRequest.toCanonical())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, customerToResponse(*customer))
}

func (u *customer) Login(c echo.Context) error {
	var customerLogin CustomerRequest

	if err := c.Bind(&customerLogin); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Message: fmt.Errorf("invalid data").Error(),
		})
	}

	token, err := u.service.Login(c.Request().Context(), customerLogin.toCanonical())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{err.Error()})
	}

	return c.JSON(http.StatusOK, TokenResponse{
		Token: token,
	})
}

func (u *customer) Bypass(c echo.Context) error {
	token, err := u.service.Bypass()
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{err.Error()})
	}
	return c.JSON(http.StatusOK, TokenResponse{
		Token: token,
	})
}
