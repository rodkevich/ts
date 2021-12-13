package ticket

import (
	"context"

	"github.com/google/uuid"

	"github.com/rodkevich/ts/ticket/internal/models"
)

type TicketsController interface {
	CreateTicket(context.Context, *models.Ticket) (*models.Ticket, error)
	GetTicket(context.Context, uuid.UUID) (*models.Ticket, error)
	ListTickets(context.Context, *models.Filter) (*models.TicketsList, error)
	UpdateTicket(context.Context, uuid.UUID) (*models.Ticket, error)
	DeleteTicket(context.Context, *models.Ticket) (*models.Ticket, error)
}

type TagController interface {
	CreateTag(context.Context, *models.Tag) (*models.Tag, error)
	GetTag(context.Context, uuid.UUID) (*models.Tag, error)
	ListTags(context.Context, *models.Tag) (*models.TagList, error)
	UpdateTag(context.Context, uuid.UUID) (*models.Tag, error)
	DeleteTag(context.Context, *models.Tag) (*models.Tag, error)
}
