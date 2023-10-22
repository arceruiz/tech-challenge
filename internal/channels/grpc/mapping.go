package grpc

import "tech-challenge/internal/canonical"

func unmarshal(customer *Customer) *canonical.Customer {
	return &canonical.Customer{
		Document: customer.Document,
		Email:    customer.Email,
		Password: customer.Password,
	}
}

func marshal(customer *canonical.Customer) *Customer {
	return &Customer{
		Email:    customer.Email,
		Document: customer.Document,
		Password: customer.Password,
	}
}
