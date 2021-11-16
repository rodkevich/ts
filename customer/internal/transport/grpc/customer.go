package grpc

import (
	"context"

	"github.com/rodkevich/ts/customer"
	"github.com/rodkevich/ts/customer/pkg/logger"
	v1 "github.com/rodkevich/ts/proto/customer/v1"
)

type UserService struct {
	v1.UnimplementedCustomerServiceServer

	customerUC customer.UsageSchema
	logger     logger.Logger
}

func (u *UserService) CreateCustomer(ctx context.Context, request *v1.CreateCustomerRequest) (*v1.CreateCustomerResponse, error) {
	panic("implement me")
}

func (u *UserService) ListCustomers(ctx context.Context, request *v1.ListCustomersRequest) (*v1.ListCustomersResponse, error) {
	panic("implement me")
}

func (u *UserService) UpdateCustomer(ctx context.Context, request *v1.UpdateCustomerRequest) (*v1.ListCustomersResponse, error) {
	panic("implement me")
}

func (u *UserService) DeleteCustomer(ctx context.Context, request *v1.DeleteCustomerRequest) (*v1.DeleteCustomerResponse, error) {
	panic("implement me")
}

func (u *UserService) Login(ctx context.Context, request *v1.LoginRequest) (*v1.LoginResponse, error) {
	panic("implement me")
}

func (u *UserService) Logout(ctx context.Context, request *v1.LogoutRequest) (*v1.LogoutResponse, error) {
	panic("implement me")
}
