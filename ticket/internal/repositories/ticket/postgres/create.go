package postgres

import (
	"context"
	"fmt"

	"github.com/rodkevich/ts/ticket/internal/models"
)

const createTicket = `
	INSERT INTO tickets
	(owner_id, name_short, name_ext, description, amount, 
	price, currency, priority, published)
	VALUES
	($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING
	id, owner_id, name_short, name_ext, description, 
	amount, price, currency, priority, published, 
	active, created_at, updated_at, deleted
	`

func (tpg *ticketPG) Create(ctx context.Context, arg *models.Ticket) (*models.Ticket, error) {
	fmt.Printf("pg Ticket Create: %+v\n", arg)

	row := tpg.db.QueryRow(
		ctx, createTicket,
		arg.OwnerID, arg.NameShort, arg.NameExt,
		arg.Description, arg.Amount, arg.Price,
		arg.Currency, arg.Priority, arg.Published,
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
