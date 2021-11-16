package customer

import (
	"context"
	"github.com/google/uuid"
	"github.com/rodkevich/ts/customer/internal/resources"
)

type UsageSchema interface {
	CreateCustomer(context.Context, *resources.Customer) (uuid.UUID, error)
	ListCustomers(context.Context, *resources.Customer) ([]*resources.Customer, error)
	UpdateCustomer(context.Context, uuid.UUID) (*resources.Customer, error)
	DeleteCustomer(context.Context, *resources.Customer) (*resources.Customer, error)

	Login(ctx context.Context, id string, password string) error
	Logout(ctx context.Context, id string) error
}
