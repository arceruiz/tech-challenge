package adapter

import (
	"client/internal/canonical"
	"fmt"

	restPorts "client/internal/rest/port"
	"client/internal/service/adapter"
	"client/internal/service/port"
	"net/http"
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
		return c.JSON(http.StatusBadRequest, fmt.Errorf("invalid data"))
	}

	customer, err := u.service.Create(canonical.Customer{
		Document:  customerRequest.Document,
		Email:     customerRequest.Email,
		Name:      customerRequest.Name,
		Password:  customerRequest.Password,
		CreatedAt: time.Now().String(),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, customer)
}

func (u *customerPort) Login(c echo.Context) error {
	var customerLogin CustomerRequest

	if err := c.Bind(&customerLogin); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("invalid data"))
	}

	token, err := u.service.Login(canonical.Customer{
		Email:    customerLogin.Email,
		Password: customerLogin.Password,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, TokenResponse{
		Token: token,
	})
}

func (u *customerPort) Bypass(c echo.Context) error {
	token, err := u.service.Bypass()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, TokenResponse{
		Token: token,
	})
}
