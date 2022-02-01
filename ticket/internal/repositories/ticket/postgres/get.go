package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/rodkevich/ts/ticket/internal/models"
	"github.com/rodkevich/ts/ticket/pkg/filter"
)

func (tpg *ticketRepositoryPG) Get(ctx context.Context, id uuid.UUID) (*models.Ticket, error) {
	const getTicket = `
	SELECT
	id, owner_id, name_short, name_ext, description,
	amount, price, currency, priority, published,
	active, created_at, updated_at, deleted
	FROM tickets
	WHERE id = $1 and tickets.deleted = 'false'
	LIMIT 1
	`
	row := tpg.db.QueryRow(
		ctx, getTicket, id,
	)
	var rtn models.Ticket
	err := row.Scan(
		&rtn.ID, &rtn.OwnerID, &rtn.NameShort, &rtn.NameExt, &rtn.Description,
		&rtn.Amount, &rtn.Price, &rtn.Currency, &rtn.Priority, &rtn.Published,
		&rtn.Active, &rtn.CreatedAt, &rtn.UpdatedAt, &rtn.Deleted,
	)

	cur := filter.EncodeCursor(rtn.CreatedAt, rtn.ID.String())
	fmt.Println("\n GET EncodeCursor", cur)

	t, u, err := filter.DecodeCursor(cur)
	fmt.Println("\n", t, u, err)

	return &rtn, err
}
