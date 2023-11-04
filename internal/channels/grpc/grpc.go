package grpc

import (
	"context"
	"net"
	"strconv"
	"tech-challenge/internal/service"

	protocol "google.golang.org/grpc"
)

type grpcServer struct {
	svc service.CustomerService
	UnimplementedCustomerServiceServer
}

func Listen(port int) error {
	server := protocol.NewServer()
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return err
	}

	RegisterCustomerServiceServer(server, &grpcServer{
		svc: service.NewCustomerService(),
	})

	return server.Serve(listener)
}

func (r *grpcServer) Get(ctx context.Context, customer *Customer) (*Customer, error) {
	request := unmarshal(customer)

	response, err := r.svc.Get(ctx, *request)
	if err != nil {
		return nil, err
	}

	return marshal(response), nil
}
