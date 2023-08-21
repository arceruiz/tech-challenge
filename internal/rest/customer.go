package rest

import (
	"fmt"
	"tech-challenge/internal/canonical"
	"tech-challenge/internal/service"

	"net/http"
	"time"

	"github.com/labstack/echo"
)

type Customer interface {
	Register(g *echo.Group)
	Create(c echo.Context) error
	Login(c echo.Context) error
}

type customer struct {
	service service.CustomerService
}

func NewCustomerChannel() Customer {
	return &customer{
		service: service.NewCustomerService(),
	}
}

func (u *customer) Register(g *echo.Group) {
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

	customer, err := u.service.Create(canonical.Customer{
		Document:  customerRequest.Document,
		Email:     customerRequest.Email,
		Name:      customerRequest.Name,
		Password:  customerRequest.Password,
		CreatedAt: time.Now().Format(time.RFC822Z),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, CustomerResponse{
		ID:       customer.Id,
		Name:     customer.Name,
		Document: customer.Document,
		Email:    customer.Email,
	})
}

func (u *customer) Login(c echo.Context) error {
	var customerLogin CustomerRequest

	if err := c.Bind(&customerLogin); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Message: fmt.Errorf("invalid data").Error(),
		})
	}

	token, err := u.service.Login(canonical.Customer{
		Email:    customerLogin.Email,
		Password: customerLogin.Password,
	})
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
