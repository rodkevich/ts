package postgres

import (
	"context"

	"github.com/google/uuid"

	"github.com/rodkevich/ts/ticket/internal/models"
)

func (tpg *ticketRepositoryPG) Update(ctx context.Context, arg models.UpdateTicketParams, id uuid.UUID) (*models.Ticket, error) {
	const updateTicket = `
	UPDATE tickets
	SET owner_id=$1, name_short=$2, name_ext=$3, description=$4, amount=$5, price=$6, currency=$7, priority=$8, published=$9, active=$10
	WHERE id = $11
	RETURNING
	id, owner_id, name_short, name_ext, description, amount, price, currency, priority, published, active, created_at, updated_at, deleted
	`
	row := tpg.db.QueryRow(
		ctx, updateTicket,
		arg.OwnerID, arg.NameShort, arg.NameExt, arg.Description,
		arg.Amount, arg.Price, arg.Currency, arg.Priority,
		arg.Published, arg.Active, id,
	)
	var rtn models.Ticket
	err := row.Scan(
		&rtn.ID, &rtn.OwnerID, &rtn.NameShort, &rtn.NameExt, &rtn.Description,
		&rtn.Amount, &rtn.Price, &rtn.Currency, &rtn.Priority, &rtn.Published,
		&rtn.Active, &rtn.CreatedAt, &rtn.UpdatedAt, &rtn.Deleted,
	)
	return &rtn, err
}
