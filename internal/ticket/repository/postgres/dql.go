package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createTicket = `-- name: CreateTicket :one
INSERT INTO tickets(owner_id, name_short, name_ext, description, amount, price,
                    currency, advantage)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id
`

type CreateTicketParams struct {
	OwnerID     uuid.UUID
	NameShort   string
	NameExt     string
	Description *string
	Amount      int32
	Price       float64
	Currency    int32
	Advantage   EnumTicketsAdvantagesType
}

func (q *Queries) CreateTicket(ctx context.Context, arg CreateTicketParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, createTicket,
		arg.OwnerID,
		arg.NameShort,
		arg.NameExt,
		arg.Description,
		arg.Amount,
		arg.Price,
		arg.Currency,
		arg.Advantage,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const deleteTicket = `-- name: DeleteTicket :exec
DELETE
FROM tickets
WHERE id = '$1'
`

func (q *Queries) DeleteTicket(ctx context.Context) error {
	_, err := q.db.Exec(ctx, deleteTicket)
	return err
}

const listTicket = `-- name: ListTicket :many
SELECT id, owner_id, name_short, name_ext, description, amount, price, currency, active, advantage, published_at, created_at, updated_at, deleted_at
FROM tickets
ORDER BY updated_at DESC
`

func (q *Queries) ListTicket(ctx context.Context) ([]Ticket, error) {
	rows, err := q.db.Query(ctx, listTicket)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Ticket
	for rows.Next() {
		var i Ticket
		if err := rows.Scan(
			&i.ID,
			&i.OwnerID,
			&i.NameShort,
			&i.NameExt,
			&i.Description,
			&i.Amount,
			&i.Price,
			&i.Currency,
			&i.Active,
			&i.Advantage,
			&i.PublishedAt,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const readTicket = `-- name: ReadTicket :one
SELECT id, owner_id, name_short, name_ext, description, amount, price, currency, active, advantage, published_at, created_at, updated_at, deleted_at
FROM tickets
WHERE id = $1
LIMIT 1
`

func (q *Queries) ReadTicket(ctx context.Context, id uuid.UUID) (Ticket, error) {
	row := q.db.QueryRow(ctx, readTicket, id)
	var i Ticket
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.NameShort,
		&i.NameExt,
		&i.Description,
		&i.Amount,
		&i.Price,
		&i.Currency,
		&i.Active,
		&i.Advantage,
		&i.PublishedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const updateTicket = `-- name: UpdateTicket :one
UPDATE tickets
SET id=$1,
    owner_id=$2,
    name_short=$3,
    name_ext=$4,
    description=$5,
    amount=$6,
    price=$7,
    currency=$8,
    active=$9,
    advantage=$10,
    published_at=$11,
    created_at=$12,
    updated_at=$13,
    deleted_at=$14
WHERE id = $1
RETURNING id
`

type UpdateTicketParams struct {
	ID          uuid.UUID
	OwnerID     uuid.UUID
	NameShort   string
	NameExt     string
	Description *string
	Amount      int32
	Price       float64
	Currency    int32
	Active      bool
	Advantage   EnumTicketsAdvantagesType
	PublishedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func (q *Queries) UpdateTicket(ctx context.Context, arg UpdateTicketParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, updateTicket,
		arg.ID,
		arg.OwnerID,
		arg.NameShort,
		arg.NameExt,
		arg.Description,
		arg.Amount,
		arg.Price,
		arg.Currency,
		arg.Active,
		arg.Advantage,
		arg.PublishedAt,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.DeletedAt,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}
