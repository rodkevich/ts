package postgres

import "github.com/jackc/pgx/v4/pgxpool"

type ticketPG struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *ticketPG {
	return &ticketPG{db: db}
}
