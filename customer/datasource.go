package customer

import (
	"context"

	"github.com/google/uuid"

	"github.com/rodkevich/ts/customer/internal/models"
)

type Proprietor interface {
	CreateCustomer(ctx context.Context, arg models.CreateCustomerParams) (models.Customer, error)
	DeleteCustomer(ctx context.Context, id uuid.UUID) error
	GetCustomer(ctx context.Context, id uuid.UUID) (models.Customer, error)
	ListCustomers(ctx context.Context) (*models.CustomersList, error)
	UpdateCustomer(ctx context.Context, arg models.UpdateCustomerParams) (models.Customer, error)
}
