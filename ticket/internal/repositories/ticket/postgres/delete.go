package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/rodkevich/ts/ticket/internal/models"
)

const deleteTicketHard = `
	DELETE FROM tickets WHERE id = ($1)
	`
const changeTicketDeletedState = `
	UPDATE tickets
	SET deleted='true'
	WHERE id = $1
	RETURNING id, owner_id, name_short, name_ext, description, amount, price, currency, priority, published, active, created_at, updated_at, deleted
	`

func (tpg *ticketRepositoryPG) Delete(ctx context.Context, id uuid.UUID, hardDelete bool) error {
	fmt.Printf("pg Delete: %+v %+v\n", id, hardDelete)

	if hardDelete {
		_, err := tpg.db.Exec(
			ctx, deleteTicketHard, id,
		)
		return err
	}

	row := tpg.db.QueryRow(
		ctx, changeTicketDeletedState, id,
	)
	var rtn models.Ticket
	err := row.Scan(
		&rtn.ID, &rtn.OwnerID, &rtn.NameShort, &rtn.NameExt,
		&rtn.Description, &rtn.Amount, &rtn.Price,
		&rtn.Currency, &rtn.Priority, &rtn.Published,
		&rtn.Active, &rtn.CreatedAt, &rtn.UpdatedAt, &rtn.Deleted,
	)
	return err

}
