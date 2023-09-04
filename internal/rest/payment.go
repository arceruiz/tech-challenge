package rest

import (
	"fmt"
	"tech-challenge/internal/service"

	"net/http"

	"github.com/labstack/echo"
)

type Payment interface {
	RegisterGroup(g *echo.Group)
	Callback(c echo.Context) error
}

type payment struct {
	service service.OrderService
}

func NewPaymentChannel() Payment {
	return &payment{
		service: service.NewOrderService(),
	}
}

func (p *payment) RegisterGroup(g *echo.Group) {
	indexPath := "/"
	g.POST(indexPath+"/callback", p.Callback)
}

func (p *payment) Callback(c echo.Context) error {

	var callback PaymentCallback
	if err := c.Bind(&callback); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Message: fmt.Errorf("invalid data").Error(),
		})
	}

	err := p.service.PaymentCallback(c.Request().Context(), callback.OrderID, callback.Status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error processing callback")
	}

	return c.JSON(http.StatusOK, nil)
}
