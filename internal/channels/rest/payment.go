package rest

import (
	"fmt"
	"tech-challenge/internal/service"

	"net/http"

	"github.com/labstack/echo/v4"
)

type Payment interface {
	RegisterGroup(g *echo.Group)
	Callback(c echo.Context) error
	GetByID(c echo.Context) error
}

type payment struct {
	orderSvc   service.OrderService
	paymentSvc service.PaymentService
}

func NewPaymentChannel() Payment {
	return &payment{
		orderSvc:   service.NewOrderService(),
		paymentSvc: service.NewPaymentService(),
	}
}

func (p *payment) RegisterGroup(g *echo.Group) {
	indexPath := "/"
	g.GET(indexPath+":id", p.GetByID)
	g.POST(indexPath+"callback", p.Callback)

}

func (p *payment) GetByID(c echo.Context) error {
	id := c.Param("id")
	if len(id) == 0 {
		return c.JSON(http.StatusBadRequest, Response{
			Message: "missing id query param",
		})
	}

	payment, err := p.paymentSvc.GetByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error processing callback")
	}

	return c.JSON(http.StatusOK, payment)
}

func (p *payment) Callback(c echo.Context) error {

	var callback PaymentCallback
	if err := c.Bind(&callback); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Message: fmt.Errorf("invalid data").Error(),
		})
	}

	err := p.orderSvc.PaymentCallback(c.Request().Context(), callback.OrderID, callback.Status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error processing callback")
	}

	return c.JSON(http.StatusOK, nil)
}
