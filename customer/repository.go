package customer

import (
	"context"
	"github.com/google/uuid"
	"github.com/rodkevich/ts/customer/internal/resources"
)

type Repository interface {
	CreateCustomer(ctx context.Context, arg resources.CreateCustomerParams) (resources.Customer, error)
	DeleteCustomer(ctx context.Context, id uuid.UUID) error
	GetCustomer(ctx context.Context, id uuid.UUID) (resources.Customer, error)
	ListCustomers(ctx context.Context) ([]resources.Customer, error)
	UpdateCustomer(ctx context.Context, arg resources.UpdateCustomerParams) (resources.Customer, error)
}
