package rest

import "tech-challenge/internal/canonical"

func (p *ProductRequest) toCanonical() canonical.Product {
	return canonical.Product{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Category:    p.Category,
		Status:      p.Status,
		ImagePath:   p.ImagePath,
	}
}

func productToResponse(p canonical.Product) ProductResponse {
	return ProductResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Category:    p.Category,
		Status:      p.Status,
		ImagePath:   p.ImagePath,
	}
}

func (c *CustomerRequest) toCanonical() canonical.Customer {
	return canonical.Customer{
		Name:     c.Name,
		Document: c.Document,
		Email:    c.Email,
		Password: c.Password,
	}
}

func customerToResponse(customer canonical.Customer) CustomerResponse {
	return CustomerResponse{
		ID:       customer.Id,
		Name:     customer.Name,
		Document: customer.Document,
		Email:    customer.Email,
	}
}

func (o *OrderRequest) toCanonical() canonical.Order {
	var orderItems []canonical.OrderItem

	for _, item := range o.OrderItems {
		orderItems = append(orderItems, item.toCanonical())
	}

	return canonical.Order{
		CustomerID: o.CustomerID,
		PaymentID:  o.PaymentID,
		Status:     o.Status,
		OrderItems: orderItems,
	}
}

func (items *OrderItem) toCanonical() canonical.OrderItem {
	return canonical.OrderItem{
		Product:  items.Product.toCanonical(),
		Quantity: items.Quantity,
	}
}
