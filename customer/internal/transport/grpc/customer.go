package grpc

import (
	"context"

	"github.com/rodkevich/ts/customer"
	"github.com/rodkevich/ts/customer/pkg/logger"
	v1 "github.com/rodkevich/ts/proto/customer/v1"
)

type CustomerService struct {
	v1.UnimplementedCustomerServiceServer

	customerUC customer.UsageSchema
	logger     logger.Logger
}

func (c *CustomerService) CreateCustomer(ctx context.Context, r *v1.CreateCustomerRequest) (*v1.CreateCustomerResponse, error) {
	panic("implement me")
}

func (c *CustomerService) ListCustomers(ctx context.Context, r *v1.ListCustomersRequest) (*v1.ListCustomersResponse, error) {
	panic("implement me")
}

func (c *CustomerService) UpdateCustomer(ctx context.Context, r *v1.UpdateCustomerRequest) (*v1.ListCustomersResponse, error) {
	panic("implement me")
}

func (c *CustomerService) DeleteCustomer(ctx context.Context, r *v1.DeleteCustomerRequest) (*v1.DeleteCustomerResponse, error) {
	panic("implement me")
}

func (c *CustomerService) Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error) {
	panic("implement me")
}

func (c *CustomerService) Logout(ctx context.Context, r *v1.LogoutRequest) (*v1.LogoutResponse, error) {
	panic("implement me")
}
