package grpc

import (
	"context"

	"github.com/google/uuid"

	"github.com/rodkevich/ts/customer"
	"github.com/rodkevich/ts/customer/internal/resources"
	"github.com/rodkevich/ts/customer/pkg/logger"
	"github.com/rodkevich/ts/customer/pkg/types"
	v1 "github.com/rodkevich/ts/customer/proto/v1"
)

type GrpcCustomerService struct {
	v1.UnimplementedCustomerServiceServer

	UseSchema customer.UsageSchema
	logger    logger.Logger
}

func (c *GrpcCustomerService) CreateCustomer(ctx context.Context, r *v1.CreateCustomerRequest) (*v1.CreateCustomerResponse, error) {

	res := resources.Customer{
		ID:        uuid.MustParse(r.Customer.GetId()),
		Type:      types.EnumCustomersType(r.Customer.GetType()),
		Status:    types.EnumCustomersStatus(r.Customer.GetStatus()),
		Login:     r.Customer.GetLogin(),
		Password:  r.Customer.GetPassword(),
		Identity:  &r.Customer.Identity,
		CreatedAt: r.Customer.GetCreatedAt().AsTime(),
		UpdatedAt: r.Customer.GetUpdatedAt().AsTime(),
		Deleted:   false,
	}

	customerId, err := c.UseSchema.CreateCustomer(ctx, &res)
	if err != nil {
		return nil, err
	}
	resp := v1.CreateCustomerResponse{CustomerId: customerId.String()}
	return &resp, nil
}

// func (c *Handler) ListCustomers(ctx context.Context, r *v1.ListCustomersRequest) (*v1.ListCustomersResponse, error) {
// 	panic("implement me")
// }
//
// func (c *Handler) UpdateCustomer(ctx context.Context, r *v1.UpdateCustomerRequest) (*v1.ListCustomersResponse, error) {
// 	panic("implement me")
// }
//
// func (c *Handler) DeleteCustomer(ctx context.Context, r *v1.DeleteCustomerRequest) (*v1.DeleteCustomerResponse, error) {
// 	panic("implement me")
// }
//
// func (c *Handler) Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error) {
// 	panic("implement me")
// }
//
// func (c *Handler) Logout(ctx context.Context, r *v1.LogoutRequest) (*v1.LogoutResponse, error) {
// 	panic("implement me")
// }
