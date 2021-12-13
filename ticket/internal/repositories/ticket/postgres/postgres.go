package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/rodkevich/ts/ticket/internal/models"
)

type ticketPG struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *ticketPG {
	return &ticketPG{db: db}
}

func (tpg *ticketPG) CreateTicket(ctx context.Context, arg *models.Ticket) (*models.Ticket, error) {
	//fmt.Printf("%+v\n", arg)
	const createTicket = `
	INSERT INTO tickets
	(owner_id, name_short, name_ext, description, amount, 
	price, currency, priority, published, active)
	VALUES
	($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING
	id, owner_id, name_short, name_ext, description, 
	amount, price, currency, priority, published, 
	active, created_at, updated_at, deleted
	`
	row := tpg.db.QueryRow(
		ctx, createTicket,
		arg.OwnerID, arg.NameShort, arg.NameExt,
		arg.Description, arg.Amount, arg.Price,
		arg.Currency, arg.Priority, arg.Published,
		arg.Active,
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

func (tpg *ticketPG) GetTicket(ctx context.Context, id uuid.UUID) (*models.Ticket, error) {
	const getTicket = `
	SELECT
	id, owner_id, name_short, name_ext, description, amount, price, currency, priority, published, active, created_at, updated_at, deleted
	FROM tickets
	WHERE id = $1
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
	return &rtn, err
}

func (tpg *ticketPG) ListTickets(ctx context.Context, f *models.Filter) (*models.TicketsList, error) {
	const listTickets = `
	SELECT
	id, owner_id, name_short, name_ext, description, amount, price, currency, priority, published, active, created_at, updated_at, deleted
	FROM tickets
	ORDER BY updated_at
	DESC
	`
	rows, err := tpg.db.Query(
		ctx, listTickets,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rtn := make([]*models.Ticket, 0)
	for rows.Next() {
		var each models.Ticket
		if err := rows.Scan(
			&each.ID, &each.OwnerID, &each.NameShort, &each.NameExt,
			&each.Description, &each.Amount, &each.Price, &each.Currency,
			&each.Priority, &each.Published, &each.Active, &each.CreatedAt,
			&each.UpdatedAt, &each.Deleted,
		); err != nil {
			return nil, err
		}
		rtn = append(rtn, &each)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &models.TicketsList{Tickets: rtn}, nil
}

func (tpg *ticketPG) UpdateTicket(ctx context.Context, arg models.UpdateTicketParams, id uuid.UUID) (*models.Ticket, error) {
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

func (tpg *ticketPG) ChangeTicketActivenessState(ctx context.Context, active bool, id uuid.UUID) (*models.Ticket, error) {
	const changeTicketActivenessState = `
	UPDATE tickets
	SET active=$1
	WHERE id = $2
	RETURNING
	id, owner_id, name_short, name_ext, description, amount, price, currency, priority, published, active, created_at, updated_at, deleted
	`
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

func (tpg *ticketPG) ChangeTicketPublishState(ctx context.Context, published bool, id uuid.UUID) (*models.Ticket, error) {
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

func (tpg *ticketPG) ChangeTicketDeletedState(ctx context.Context, deleted bool, id uuid.UUID) (*models.Ticket, error) {
	const changeTicketDeletedState = `
	UPDATE tickets
	SET deleted=$1
	WHERE id = $2
	RETURNING id, owner_id, name_short, name_ext, description, amount, price, currency, priority, published, active, created_at, updated_at, deleted
	`
	row := tpg.db.QueryRow(
		ctx, changeTicketDeletedState, deleted, id,
	)
	var rtn models.Ticket
	err := row.Scan(
		&rtn.ID, &rtn.OwnerID, &rtn.NameShort, &rtn.NameExt,
		&rtn.Description, &rtn.Amount, &rtn.Price,
		&rtn.Currency, &rtn.Priority, &rtn.Published,
		&rtn.Active, &rtn.CreatedAt, &rtn.UpdatedAt, &rtn.Deleted,
	)
	return &rtn, err
}

func (tpg *ticketPG) DeleteTicket(ctx context.Context, id uuid.UUID) error {
	const deleteTicket = `
	DELETE FROM tickets
	WHERE id = $1
	`
	_, err := tpg.db.Exec(
		ctx, deleteTicket, id,
	)
	return err
}
