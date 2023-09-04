package repository

import "tech-challenge/internal/canonical"

func (order *Order) toCanonical() canonical.Order {
	return canonical.Order{
		ID:         order.ID,
		CustomerID: order.CustomerID,
		Status:     canonical.OrderStatus(order.Status),
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
		Total:      order.Total,
		Payment: &canonical.Payment{
			ID:          *order.PaymentID,
			PaymentType: *order.PaymentID,
			CreatedAt:   order.paymentCreatedat,
		},
	}
}
