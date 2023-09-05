package rest

import (
	"context"
	"tech-challenge/internal/service"

	"net/http"

	"github.com/labstack/echo/v4"
)

type Product interface {
	RegisterGroup(g *echo.Group)
	Get(c echo.Context) error
	Add(c echo.Context) error
	Update(c echo.Context) error
	Remove(c echo.Context) error
}

type product struct {
	service service.ProductService
}

func NewProductChannel() Product {
	return &product{
		service: service.NewProductService(),
	}
}

func (p *product) RegisterGroup(g *echo.Group) {
	indexPath := "/"
	g.GET(indexPath+":id", p.Get)
	g.GET("", p.GetAll)
	g.POST(indexPath, p.Add)
	g.PUT(indexPath+":id", p.Update)
	g.DELETE(indexPath+":id", p.Remove)
}

func (p *product) GetAll(ctx echo.Context) error {
	category := ctx.QueryParam("category")

	response, err := p.get(ctx.Request().Context(), "", category)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return ctx.JSON(http.StatusOK, response)
}

func (p *product) Get(ctx echo.Context) error {
	productID := ctx.QueryParam("id")

	response, err := p.get(ctx.Request().Context(), productID, "")
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	if len(response) == 0 {
		return ctx.JSON(http.StatusNotFound, nil)
	}

	return ctx.JSON(http.StatusOK, response[0])
}

func (p *product) get(ctx context.Context, productID string, category string) ([]ProductResponse, error) {

	var response []ProductResponse

	if productID != "" {
		product, err := p.service.GetByID(ctx, productID)
		if err != nil {
			return nil, err
		}
		return []ProductResponse{productToResponse(*product)}, nil
	}

	if category != "" {
		products, err := p.service.GetByCategory(ctx, category)
		if err != nil {
			return nil, err
		}

		for _, product := range products {
			response = append(response, productToResponse(product))
		}
		return response, nil
	}

	products, err := p.service.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	for _, product := range products {
		response = append(response, productToResponse(product))
	}

	return response, nil
}

func (p *product) Add(c echo.Context) error {
	var newProduct ProductRequest
	err := c.Bind(&newProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	insertedId, err := p.service.CreateProduct(c.Request().Context(), newProduct.toCanonical())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, ProductResponse{
		ID: insertedId,
	})
}

func (p *product) Update(c echo.Context) error {
	productID := c.Param("id")

	var updatedProduct ProductRequest
	err := c.Bind(&updatedProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	err = p.service.UpdateProduct(c.Request().Context(), productID, updatedProduct.toCanonical())
	if err != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}

	return c.JSON(http.StatusOK, nil)
}

func (p *product) Remove(c echo.Context) error {
	productID := c.Param("id")

	err := p.service.Remove(c.Request().Context(), productID)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}

	return c.NoContent(http.StatusOK)
}
