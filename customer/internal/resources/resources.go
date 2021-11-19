package resources

import (
	"time"

	"github.com/google/uuid"

	"github.com/rodkevich/ts/customer/pkg/types"
)

type CreateCustomerParams struct {
	Type     types.EnumCustomersType `json:"type,omitempty"`
	Login    string                  `json:"login,omitempty"`
	Password string                  `json:"password,omitempty"`
	Identity *string                 `json:"identity"`
}

type UpdateCustomerParams struct {
	ID        uuid.UUID                 `json:"id,omitempty"`
	Type      types.EnumCustomersType   `json:"type,omitempty"`
	Status    types.EnumCustomersStatus `json:"status,omitempty"`
	Login     string                    `json:"login,omitempty"`
	Password  string                    `json:"password,omitempty"`
	Identity  *string                   `json:"identity"`
	CreatedAt time.Time                 `json:"created_at"`
	UpdatedAt time.Time                 `json:"updated_at"`
	Deleted   bool                      `json:"deleted,omitempty"`
}

type Customer struct {
	ID        uuid.UUID                 `json:"id,omitempty"`
	Type      types.EnumCustomersType   `json:"type,omitempty"`
	Status    types.EnumCustomersStatus `json:"status,omitempty"`
	Login     string                    `json:"login,omitempty"`
	Password  string                    `json:"password,omitempty"`
	Identity  *string                   `json:"identity"`
	CreatedAt time.Time                 `json:"created_at"`
	UpdatedAt time.Time                 `json:"updated_at"`
	Deleted   bool                      `json:"deleted,omitempty"`
}
