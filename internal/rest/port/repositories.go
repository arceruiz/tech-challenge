package port

import "github.com/labstack/echo"

type ProductPort interface {
	Register(g *echo.Group)
	Get(c echo.Context) error
	Add(c echo.Context) error
	Update(c echo.Context) error
	Remove(c echo.Context) error
}

type CustomerPort interface {
	Register(g *echo.Group)
	Create(c echo.Context) error
	Login(c echo.Context) error
}

type OrderPort interface {
	Register(g *echo.Group)
	Get(c echo.Context) error
	CheckoutOrder(c echo.Context) error
}
