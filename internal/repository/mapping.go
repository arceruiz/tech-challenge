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
	}
}

func (order *OrderJoinPayment) toCanonical() canonical.Order {
	var payment *canonical.Payment
	if order.PaymentID != nil {
		payment = &canonical.Payment{
			ID:          *order.PaymentID,
			PaymentType: *order.PaymentType,
			CreatedAt:   order.PaymentCreatedat,
			Status:      canonical.PaymentStatus(*order.PaymentStatus),
		}
	}
	return canonical.Order{
		ID:         order.ID,
		CustomerID: order.CustomerID,
		Status:     canonical.OrderStatus(order.Status),
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
		Total:      order.Total,
		Payment:    payment,
	}
}
