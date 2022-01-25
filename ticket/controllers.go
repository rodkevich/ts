package ticket

import (
	"context"

	"github.com/google/uuid"

	"github.com/rodkevich/ts/ticket/internal/models"
)

// TicketsController ..
type TicketsController interface {
	CreateTicket(ctx context.Context, prototype *models.Ticket) (*models.Ticket, error)
	GetTicket(ctx context.Context, id uuid.UUID) (*models.Ticket, error)
	ListTickets(ctx context.Context, filter *models.TicketFilter) (*models.TicketsList, error)
	UpdateTicket(ctx context.Context, target *models.Ticket) (*models.Ticket, error)
	DeleteTicket(ctx context.Context, id uuid.UUID, hardDelete bool) (*models.Ticket, error)
}

// TagsController ...
type TagsController interface {
	CreateTag(context.Context, *models.Tag) (*models.Tag, error)
	GetTag(context.Context, uuid.UUID) (*models.Tag, error)
	ListTags(context.Context, *models.TicketFilter) (*models.TagList, error)
	UpdateTag(context.Context, *models.Tag) (*models.Tag, error)
	DeleteTag(context.Context, uuid.UUID) (*models.Tag, error)
}
