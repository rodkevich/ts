package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/rodkevich/ts/ticket/internal/models"
)

func (tpg *ticketPG) ChangePublishState(ctx context.Context, published bool, id uuid.UUID) (*models.Ticket, error) {
	const changeTicketPublishState = `
	UPDATE tickets
	SET published=$1
	WHERE id = $2
	RETURNING
	id, owner_id, name_short, name_ext, description, amount, price, currency, priority, published, active, created_at, updated_at, deleted
`
	row := tpg.db.QueryRow(
		ctx, changeTicketPublishState, published, id,
	)
	var rtn models.Ticket
	err := row.Scan(
		&rtn.ID, &rtn.OwnerID, &rtn.NameShort, &rtn.NameExt,
		&rtn.Description, &rtn.Amount, &rtn.Price, &rtn.Currency,
		&rtn.Priority, &rtn.Published, &rtn.Active, &rtn.CreatedAt,
		&rtn.UpdatedAt, &rtn.Deleted,
	)
	return &rtn, err
}
