package resources

import (
	"time"

	"github.com/google/uuid"

	"github.com/rodkevich/ts/customer/pkg/types"
)

type CreateCustomerParams struct {
	Type     types.EnumCustomersType
	Login    string
	Password string
	Identity types.NullString
}

type UpdateCustomerParams struct {
	ID        uuid.UUID
	Type      types.EnumCustomersType
	Status    types.EnumCustomersStatus
	Login     string
	Password  string
	Identity  types.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
	Deleted   bool
}

type Customer struct {
	ID        uuid.UUID
	Type      types.EnumCustomersType
	Status    types.EnumCustomersStatus
	Login     string
	Password  string
	Identity  types.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
	Deleted   bool
}
