package customer

import (
	"context"

	"github.com/google/uuid"

	"github.com/rodkevich/ts/customer/internal/models"
)

type Invoker interface {
	CreateCustomer(context.Context, *models.Customer) (uuid.UUID, error)
	ListCustomers(context.Context, *models.Customer) (*models.CustomersList, error)
	UpdateCustomer(context.Context, uuid.UUID) (*models.Customer, error)
	DeleteCustomer(context.Context, *models.Customer) (*models.Customer, error)

	Login(ctx context.Context, id string, password string) error
	Logout(ctx context.Context, id string) error
}
