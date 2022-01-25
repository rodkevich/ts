package postgres

import (
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/rodkevich/ts/ticket"
)

type ticketRepositoryPG struct {
	db *pgxpool.Pool
}

// New ...
func New(db *pgxpool.Pool) ticket.TicketsRepositoryIFace {
	return &ticketRepositoryPG{db: db}
}
