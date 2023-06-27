package port

import (
	"client/internal/canonical"

	"client/internal/service/port"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type CustomerPort struct {
	service port.CustomerService
}

func NewUserPort(userService port.CustomerService) *CustomerPort {
	return &CustomerPort{
		service: userService,
	}
}

func (u *CustomerPort) Create(g *echo.Group) {
	g.POST("/create", u.create)
	g.POST("/login", u.login)
	g.POST("/bypass", echo.MethodNotAllowedHandler)
}

func (u *CustomerPort) create(c echo.Context) error {
	var customerRequest CustomerRequest

	if err := c.Bind(&customerRequest); err != nil {
		return c.JSON(400, "dados incorretos")
	}

	customer, err := u.service.Create(canonical.Customer{
		Document:  customerRequest.Document,
		Email:     customerRequest.Email,
		Name:      customerRequest.Name,
		Password:  customerRequest.Password,
		CreatedAt: time.Now().String(),
	})

	if err != nil {
		return c.JSON(http.StatusBadRequest, "implementar um map pra erros, algo assim")
	}

	return c.JSON(http.StatusOK, customer)
}

func (u *CustomerPort) login(c echo.Context) error {
	var customerLogin CustomerRequest

	if err := c.Bind(&customerLogin); err != nil {
		return c.JSON(400, "Login ou senha incorretos")
	}

	token, err := u.service.Login(canonical.Customer{
		Email:    customerLogin.Email,
		Password: customerLogin.Password,
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, TokenResponse{
		Token: token,
	})
}
