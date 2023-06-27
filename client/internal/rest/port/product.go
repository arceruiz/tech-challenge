package port

import (
	"client/internal/canonical"
	"client/internal/service/port"
	"net/http"

	"github.com/labstack/echo"
)

type ProductPort struct {
	service port.ProductService
}

func NewProductPort(productService port.ProductService) *ProductPort {
	return &ProductPort{
		service: productService,
	}
}

func (p *ProductPort) Register(g *echo.Group) {
	indexPath := ""
	g.GET(indexPath, p.Get)  // localhost:3001/api/product
	g.POST(indexPath, p.Add) // localhost:3001/api/product
	g.PUT(indexPath+"/:id", p.Update)
	g.DELETE(indexPath+"/:id", p.Remove)
}

func (p *ProductPort) Get(c echo.Context) error {
	productID := c.QueryParam("id")
	category := c.QueryParam("category")

	if productID != "" {
		product, err := p.service.GetByID(productID)
		if err != nil {
			return c.JSON(http.StatusNotFound, "Product not found")
		}
		return c.JSON(http.StatusOK, product)
	}

	if category != "" {
		products, err := p.service.GetByCategory(category)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, products)
	}

	products, err := p.service.GetProducts()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, products)
}

func (p *ProductPort) Add(c echo.Context) error {
	var newProduct canonical.Product
	err := c.Bind(&newProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	addedProduct, err := p.service.CreateProduct(newProduct)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, addedProduct)
}

func (p *ProductPort) Update(c echo.Context) error {
	productID := c.Param("id")

	var updatedProduct canonical.Product
	err := c.Bind(&updatedProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	updatedProduct, err = p.service.UpdateProduct(productID, updatedProduct)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}

	return c.JSON(http.StatusOK, updatedProduct)
}

func (p *ProductPort) Remove(c echo.Context) error {
	productID := c.Param("id")

	err := p.service.Remove(productID)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}

	return c.NoContent(http.StatusOK)
}
