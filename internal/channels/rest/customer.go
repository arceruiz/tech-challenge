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
