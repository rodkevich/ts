package ticket

import (
	"context"

	"github.com/google/uuid"

	"github.com/rodkevich/ts/ticket/internal/models"
)

// TODO: https://jsonapi.org/format/#fetching-sparse-fieldsets
// GET /articles?include=author&fields[articles]=title,body&fields[people]=name

// TicketsRepositoryIFace ...
type TicketsRepositoryIFace interface {
	Create(ctx context.Context, arg *models.Ticket) (*models.Ticket, error)
	Get(ctx context.Context, id uuid.UUID) (*models.Ticket, error)
	List(ctx context.Context, tf models.TicketFilter) (*models.TicketsList, error)
	Update(ctx context.Context, arg models.UpdateTicketParams, id uuid.UUID) (*models.Ticket, error)
	Delete(ctx context.Context, id uuid.UUID, hardDelete bool) error
	Search(ctx context.Context, id *uuid.UUID, f *models.TicketFilter) (*models.TicketsList, *uuid.UUID, error)
	ChangeActivenessState(ctx context.Context, active bool, id uuid.UUID) (*models.Ticket, error)
	ChangePublishState(ctx context.Context, published bool, id uuid.UUID) (*models.Ticket, error)
}

// TagsProprietor ...
type TagsProprietor interface {
	Create(ctx context.Context, name string, description *string) (models.Tag, error)
	Get(ctx context.Context, id uuid.UUID) (models.Tag, error)
	List(ctx context.Context) (*models.TagList, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
