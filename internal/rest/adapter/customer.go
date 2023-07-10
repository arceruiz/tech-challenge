package adapter

import (
	"fmt"
	"tech-challenge/internal/canonical"

	"net/http"
	restPorts "tech-challenge/internal/rest/port"
	"tech-challenge/internal/service/adapter"
	"tech-challenge/internal/service/port"
	"time"

	"github.com/labstack/echo"
)

type customerPort struct {
	service port.CustomerService
}

func NewCustomerPort() restPorts.CustomerPort {
	return &customerPort{
		service: adapter.NewCustomerService(),
	}
}

func (u *customerPort) Register(g *echo.Group) {
	g.POST("/create", u.Create)
	g.POST("/login", u.Login)
	g.POST("/bypass", u.Bypass)
}

func (u *customerPort) Create(c echo.Context) error {
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

func (u *customerPort) Login(c echo.Context) error {
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

func (u *customerPort) Bypass(c echo.Context) error {
	token, err := u.service.Bypass()
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{err.Error()})
	}
	return c.JSON(http.StatusOK, TokenResponse{
		Token: token,
	})
}
