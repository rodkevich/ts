package usage

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/rodkevich/ts/customer"
	"github.com/rodkevich/ts/customer/internal/resources"
	"github.com/rodkevich/ts/customer/pkg/logger"
)

type customerUseCase struct {
	customerPGRepo customer.Repository
	log            logger.Logger
}

func New(customerPGRepo customer.Repository, log logger.Logger) *customerUseCase {
	return &customerUseCase{customerPGRepo: customerPGRepo, log: log}
}

func (c customerUseCase) CreateCustomer(ctx context.Context, r *resources.Customer) (uuid uuid.UUID, err error) {
	arg := resources.CreateCustomerParams{
		Type:     r.Type,
		Login:    r.Login,
		Password: r.Password,
		Identity: r.Identity,
	}

	createCustomer, err := c.customerPGRepo.CreateCustomer(ctx, arg)
	if err != nil {
		_ = fmt.Errorf("%w", err)
		return
	}

	return createCustomer.ID, nil
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
