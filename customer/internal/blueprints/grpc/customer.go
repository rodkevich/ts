package grpc

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/rodkevich/ts/customer"
	"github.com/rodkevich/ts/customer/internal/models"
	"github.com/rodkevich/ts/customer/pkg/logger"
	"github.com/rodkevich/ts/customer/pkg/types"
	v1 "github.com/rodkevich/ts/customer/proto/v1"
)

type CustomerGrpcService struct {
	v1.UnimplementedCustomerServiceServer

	logger    logger.Logger
	useSchema customer.Invoker
}

func NewCustomerGrpcService(logger logger.Logger, useSchema customer.Invoker) *CustomerGrpcService {
	return &CustomerGrpcService{logger: logger, useSchema: useSchema}
}

func (s *CustomerGrpcService) ListCustomers(ctx context.Context, request *v1.ListCustomersRequest) (*v1.ListCustomersResponse, error) {
	c, err := s.useSchema.ListCustomers(ctx, nil)
	if err != nil {
		s.logger.Errorf("useSchema.ListCustomers: %v", err)
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("%s: %v", "CustomerService.ListCustomers:", err))
	}
	return &v1.ListCustomersResponse{Customers: c.ToProto()}, nil
}

func (s *CustomerGrpcService) CreateCustomer(ctx context.Context, r *v1.CreateCustomerRequest) (*v1.CreateCustomerResponse, error) {
	// TODO: check passwd
	res := models.Customer{
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

	customerId, err := s.useSchema.CreateCustomer(ctx, &res)
	if err != nil {
		s.logger.Errorf("useSchema.CreateCustomer: %v", err)
		return nil, status.Errorf(codes.AlreadyExists, fmt.Sprintf("%s: %v", "CustomerService.CreateCustomer:", err))
	}
	resp := v1.CreateCustomerResponse{CustomerId: customerId.String()}
	return &resp, nil
}

func (s *CustomerGrpcService) Login(ctx context.Context, request *v1.LoginRequest) (*v1.LoginResponse, error) {
	panic("implement me")
}

func (s *CustomerGrpcService) Logout(ctx context.Context, request *v1.LogoutRequest) (*v1.LogoutResponse, error) {
	panic("implement me")
}
