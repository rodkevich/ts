package ticket

import (
	"context"

	"github.com/google/uuid"

	"github.com/rodkevich/ts/ticket/internal/models"
)

// TODO: https://jsonapi.org/format/#fetching-sparse-fieldsets
// GET /articles?include=author&fields[articles]=title,body&fields[people]=name

type TicketsProprietor interface {
	CreateTicket(ctx context.Context, arg models.CreateTicketParams) (*models.Ticket, error)
	GetTicket(ctx context.Context, id uuid.UUID) (*models.Ticket, error)
	ListTickets(ctx context.Context) (*models.TicketsList, error)
	UpdateTicket(ctx context.Context, arg models.UpdateTicketParams, id uuid.UUID) (*models.Ticket, error)

	ChangeTicketActivenessState(ctx context.Context, active bool, id uuid.UUID) (*models.Ticket, error)
	ChangeTicketPublishState(ctx context.Context, published bool, id uuid.UUID) (*models.Ticket, error)
	ChangeTicketDeletedState(ctx context.Context, deleted bool, id uuid.UUID) (*models.Ticket, error)

	DeleteTicket(ctx context.Context, id uuid.UUID) error
}
