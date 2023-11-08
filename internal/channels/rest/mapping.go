package rest

import "tech-challenge/internal/canonical"

func (p *ProductRequest) toCanonical() canonical.Product {
	return canonical.Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Category:    p.Category,
		Status:      canonical.BaseStatus(p.Status),
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
		Status:      int(p.Status),
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
		Payment: &canonical.Payment{
			ID: o.PaymentID,
		},
		OrderItems: orderItems,
	}
}

func orderToResponse(order canonical.Order) OrderResponse {
	var productsList []ProductResponse

	for _, product := range order.OrderItems {
		productsList = append(productsList, productToResponse(product.Product))
	}

	return OrderResponse{
		ID:          order.ID,
		CustomerID:  order.CustomerID,
		Status:      int(order.Status),
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
		Products:    productsList,
		PaymentRest: paymentToResponse(order.Payment),
	}
}

func (items *OrderItem) toCanonical() canonical.OrderItem {
	return canonical.OrderItem{
		Product:  items.Product.toCanonical(),
		Quantity: items.Quantity,
	}
}

func paymentToResponse(payment *canonical.Payment) *PaymentRest {
	if payment == nil {
		return nil
	}

	return &PaymentRest{
		ID:          payment.ID,
		PaymentType: payment.PaymentType,
		CreatedAt:   payment.CreatedAt,
		Status:      int(payment.Status),
	}
}

func (pr *PaymentRest) toCanonical() canonical.Payment {
	return canonical.Payment{
		PaymentType: pr.PaymentType,
	}
}
