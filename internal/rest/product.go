package rest

import (
	"tech-challenge/internal/service"

	"net/http"

	"github.com/labstack/echo"
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
	indexPath := ""
	g.GET(indexPath, p.Get)
	g.POST(indexPath, p.Add)
	g.PUT(indexPath+"/:id", p.Update)
	g.DELETE(indexPath+"/:id", p.Remove)
}

func (p *product) Get(c echo.Context) error {
	productID := c.QueryParam("id")
	category := c.QueryParam("category")
	var response []ProductResponse

	if productID != "" {
		product, err := p.service.GetByID(productID)
		if err != nil {
			return c.JSON(http.StatusNotFound, "Product not found")
		}
		return c.JSON(http.StatusOK, productToResponse(*product))
	}

	if category != "" {
		products, err := p.service.GetByCategory(category)
		if err != nil {
			return err
		}

		for _, product := range products {
			response = append(response, productToResponse(product))
		}
		return c.JSON(http.StatusOK, response)
	}

	products, err := p.service.GetProducts()
	if err != nil {
		return err
	}

	for _, product := range products {
		response = append(response, productToResponse(product))
	}

	return c.JSON(http.StatusOK, response)
}

func (p *product) Add(c echo.Context) error {
	var newProduct ProductRequest
	err := c.Bind(&newProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	err = p.service.CreateProduct(newProduct.toCanonical())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, nil)
}

func (p *product) Update(c echo.Context) error {
	productID := c.Param("id")

	var updatedProduct ProductRequest
	err := c.Bind(&updatedProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	err = p.service.UpdateProduct(productID, updatedProduct.toCanonical())
	if err != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}

	return c.JSON(http.StatusOK, nil)
}

func (p *product) Remove(c echo.Context) error {
	productID := c.Param("id")

	err := p.service.Remove(productID)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}

	return c.NoContent(http.StatusOK)
}
