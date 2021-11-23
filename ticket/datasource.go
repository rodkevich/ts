package ticket

import (
	"context"

	"github.com/google/uuid"

	"github.com/rodkevich/ts/ticket/internal/models"
)

type ProprietorTickets interface {
	CreateTicket(ctx context.Context, arg models.CreateTicketParams) (*models.Ticket, error)
	GetTicket(ctx context.Context, id uuid.UUID) (*models.Ticket, error)
	ListTickets(ctx context.Context) (*models.TicketsList, error)
	UpdateTicket(ctx context.Context, arg models.UpdateTicketParams) (*models.Ticket, error)

	ChangeTicketActivenessState(ctx context.Context, active bool) (*models.Ticket, error)
	ChangeTicketPublishState(ctx context.Context, published bool) (*models.Ticket, error)

	MarkTicketAsDeleted(ctx context.Context, deleted bool) (*models.Ticket, error)
	DeleteTicket(ctx context.Context, id uuid.UUID) error
}
