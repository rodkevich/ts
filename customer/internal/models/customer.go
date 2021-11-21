package models

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	v1 "github.com/rodkevich/ts/customer/proto/v1"

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

// ToProto ..
func (c *Customer) ToProto() *v1.Customer {
	return &v1.Customer{
		Id:     c.ID.String(),
		Type:   "",
		Status: "",
		// Status:    string(c.Type),
		Login:     c.Login,
		Password:  "",
		Identity:  "",
		CreatedAt: timestamppb.New(c.CreatedAt),
		UpdatedAt: nil,
		Deleted:   false,
	}
}

type CustomersList struct {
	// TotalCount int      `json:"totalCount"`
	// TotalPages int      `json:"totalPages"`
	// Page       int      `json:"page"`
	// Size       int      `json:"size"`
	// HasMore    bool     `json:"hasMore"`
	Customers []*Customer `json:"comments"`
}

// ToProto ..
func (c *CustomersList) ToProto() []*v1.Customer {
	customersList := make([]*v1.Customer, 0, len(c.Customers))
	for _, customer := range c.Customers {
		customersList = append(customersList, customer.ToProto())
	}
	return customersList
}
