package repository

import (
	"context"

	"github.com/google/uuid"

	"github.com/rodkevich/ts/api/internal/models"
)

type TicketRedisRepo interface {
	GetTicketByID(ctx context.Context, id uuid.UUID) (*models.Ticket, error)
	SetTicket(ctx context.Context, t *models.Ticket) error
	DeleteTicket(ctx context.Context, id uuid.UUID) error
}
