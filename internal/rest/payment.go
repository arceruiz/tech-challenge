package rest

import (
	"tech-challenge/internal/service"

	"net/http"

	"github.com/labstack/echo"
)

type Payment interface {
	RegisterGroup(g *echo.Group)
	Callback(c echo.Context) error
}

type payment struct {
	service service.PaymentService
}

func NewPaymentChannel() Payment {
	return &payment{
		service: service.NewPaymentService(),
	}
}

func (p *payment) RegisterGroup(g *echo.Group) {
	indexPath := "/"
	g.POST(indexPath+"/callback", p.Callback)
}

func (p *payment) Callback(c echo.Context) error {
	//need to implement
	var response []ProductResponse

	products, err := p.service.GetProducts(c.Request().Context())
	if err != nil {
		return err
	}

	for _, product := range products {
		response = append(response, productToResponse(product))
	}

	return c.JSON(http.StatusOK, response)
}