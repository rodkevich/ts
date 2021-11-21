package controllers

import (
	"context"
	"github.com/google/uuid"
	"github.com/rodkevich/ts/customer"
	"github.com/rodkevich/ts/customer/internal/models"
	"github.com/rodkevich/ts/customer/pkg/logger"
)

type actions struct {
	customerPGRepo customer.Proprietor
	logger         logger.Logger
}

func NewCustomerController(log logger.Logger, customerRepo customer.Proprietor) *actions {
	return &actions{customerPGRepo: customerRepo, logger: log}
}

func (a *actions) ListCustomers(ctx context.Context, m *models.Customer) (*models.CustomersList, error) {
	// TODO: handle m *models.Customer
	customers, err := a.customerPGRepo.ListCustomers(ctx)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (a *actions) UpdateCustomer(ctx context.Context, u uuid.UUID) (*models.Customer, error) {
	panic("implement me")
}

func (a *actions) DeleteCustomer(ctx context.Context, m *models.Customer) (*models.Customer, error) {
	panic("implement me")
}

func (a *actions) Login(ctx context.Context, id string, password string) error {
	panic("implement me")
}

func (a *actions) Logout(ctx context.Context, id string) error {
	panic("implement me")
}

func (a actions) CreateCustomer(ctx context.Context, m *models.Customer) (uuid uuid.UUID, err error) {
	arg := models.CreateCustomerParams{
		Type:     m.Type,
		Login:    m.Login,
		Password: m.Password,
		Identity: m.Identity,
	}

	createCustomer, err := a.customerPGRepo.CreateCustomer(ctx, arg)
	if err != nil {
		a.logger.Debug("CreateCustomer controller:", err)
		return
	}

	return createCustomer.ID, nil
}
