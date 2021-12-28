package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/rodkevich/ts/ticket/internal/models"
)

const changeTicketActivenessState = `
	UPDATE tickets
	SET active=$1
	WHERE id = $2
	RETURNING
	id, owner_id, name_short, name_ext, description, amount, price, currency, priority, published, active, created_at, updated_at, deleted
	`

func (tpg *ticketPG) ChangeActivenessState(ctx context.Context, active bool, id uuid.UUID) (*models.Ticket, error) {

	row := tpg.db.QueryRow(
		ctx, changeTicketActivenessState, active, id,
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
