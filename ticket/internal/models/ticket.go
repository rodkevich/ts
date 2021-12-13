package models

import (
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/rodkevich/ts/ticket/pkg/types"
	"github.com/rodkevich/ts/ticket/proto/ticket/v1"
)

type (
	Ticket struct {
		ID          uuid.UUID
		OwnerID     uuid.UUID
		NameShort   string
		NameExt     *string
		Description *string
		Amount      int32
		Price       float64
		Currency    int32
		Priority    types.EnumTicketsPriority
		Published   bool
		Active      bool
		CreatedAt   time.Time
		UpdatedAt   time.Time
		Deleted     bool
	}

	CreateTicketParams struct {
		OwnerID     uuid.UUID
		NameShort   string
		NameExt     *string
		Description *string
		Amount      int32
		Price       float64
		Currency    int32
		Priority    types.EnumTicketsPriority
		Published   bool
	}

	UpdateTicketParams struct {
		OwnerID     uuid.UUID
		NameShort   string
		NameExt     *string
		Description *string
		Amount      int32
		Price       float64
		Currency    int32
		Priority    types.EnumTicketsPriority
		Active      bool
		Published   bool
	}

	TicketsList struct {
		Tickets []*Ticket `json:"tickets"`
	}
)

// ToProto ..
func (t *Ticket) ToProto() *v1.Ticket {
	return &v1.Ticket{
		Id:          t.ID.String(),
		OwnerId:     t.OwnerID.String(),
		NameShort:   t.NameShort,
		NameExt:     *t.NameExt,
		Description: *t.Description,
		Amount:      t.Amount,
		Price:       t.Price,
		Currency:    t.Currency,
		Priority:    string(t.Priority),
		Published:   t.Published,
		Active:      t.Active,
		CreatedAt:   timestamppb.New(t.CreatedAt),
		UpdatedAt:   timestamppb.New(t.UpdatedAt),
		Deleted:     t.Deleted,
	}
}

// ToProto ..
func (tl *TicketsList) ToProto() []*v1.Ticket {
	customersList := make([]*v1.Ticket, 0, len(tl.Tickets))
	for _, ticket := range tl.Tickets {
		customersList = append(customersList, ticket.ToProto())
	}
	return customersList
}
