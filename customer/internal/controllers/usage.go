package controllers

import (
	"context"
	"github.com/google/uuid"
	"github.com/rodkevich/ts/customer"
	"github.com/rodkevich/ts/customer/internal/models"
	"github.com/rodkevich/ts/customer/pkg/logger"
)

type controller struct {
	customerPGRepo customer.ProprietorCustomers
	logger         logger.Logger
}

func NewCustomerController(log logger.Logger, customerRepo customer.ProprietorCustomers) *controller {
	return &controller{customerPGRepo: customerRepo, logger: log}
}

func (c controller) CreateCustomer(ctx context.Context, m *models.Customer) (uuid *models.Customer, err error) {
	arg := models.CreateCustomerParams{
		Type:     m.Type,
		Login:    m.Login,
		Password: m.Password,
		Identity: m.Identity,
	}

	createCustomer, err := c.customerPGRepo.CreateCustomer(ctx, arg)
	if err != nil {
		c.logger.Debug("CreateCustomer controller:", err)
		return
	}

	return createCustomer, nil
}

func (c *controller) GetCustomer(ctx context.Context, u uuid.UUID) (*models.Customer, error) {
	panic("not implemented")
}

func (c *controller) ListCustomers(ctx context.Context, m *models.Customer) (*models.CustomersList, error) {
	// TODO: handle m *models.Customer
	customers, err := c.customerPGRepo.ListCustomers(ctx)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (c *controller) UpdateCustomer(ctx context.Context, u uuid.UUID) (*models.Customer, error) {
	panic("implement me")
}

func (c *controller) DeleteCustomer(ctx context.Context, m *models.Customer) (*models.Customer, error) {
	panic("implement me")
}

func (c *controller) Login(ctx context.Context, id string, password string) error {
	panic("implement me")
}

func (c *controller) Logout(ctx context.Context, id string) error {
	panic("implement me")
}
