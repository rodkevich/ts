package ticket

import (
	"context"

	"github.com/google/uuid"

	"github.com/rodkevich/ts/ticket/internal/models"
)

type TicketsController interface {
	CreateTicket(context.Context, *models.Ticket) (*models.Ticket, error)
	GetTicket(context.Context, uuid.UUID) (*models.Ticket, error)
	ListTickets(context.Context, *models.TicketFilter) (*models.TicketsList, error)
	UpdateTicket(context.Context, *models.Ticket) (*models.Ticket, error)
	DeleteTicket(ctx context.Context, id uuid.UUID, hardDelete bool) (*models.Ticket, error)
}

type TagController interface {
	CreateTag(context.Context, *models.Tag) (*models.Tag, error)
	GetTag(context.Context, uuid.UUID) (*models.Tag, error)
	ListTags(context.Context, *models.TicketFilter) (*models.TagList, error)
	UpdateTag(context.Context, *models.Tag) (*models.Tag, error)
	DeleteTag(context.Context, uuid.UUID) (*models.Tag, error)
}
