package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/rodkevich/ts/customer"
	"github.com/rodkevich/ts/customer/internal/resources"
	"github.com/rodkevich/ts/customer/pkg/logger"
)

type customerUseCase struct {
	customerPGRepo customer.Repository
	log            logger.Logger
}

func (c customerUseCase) CreateCustomer(ctx context.Context, r *resources.Customer) (uuid.UUID, error) {
	panic("implement me")
}

func (c customerUseCase) ListCustomers(ctx context.Context, r *resources.Customer) ([]*resources.Customer, error) {
	panic("implement me")
}

func (c customerUseCase) UpdateCustomer(ctx context.Context, uuid uuid.UUID) (*resources.Customer, error) {
	panic("implement me")
}

func (c customerUseCase) DeleteCustomer(ctx context.Context, r *resources.Customer) (*resources.Customer, error) {
	panic("implement me")
}

func (c customerUseCase) Login(ctx context.Context, id string, password string) error {
	panic("implement me")
}

func (c customerUseCase) Logout(ctx context.Context, id string) error {
	panic("implement me")
}
