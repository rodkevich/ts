package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/google/uuid"
	"github.com/rodkevich/ts/customer/internal/resources"
)

type consumerPGRepository struct {
	db *pgxpool.Pool
}

func NewConsumerPGRepository(db *pgxpool.Pool) *consumerPGRepository {
	return &consumerPGRepository{db: db}
}

const (
	createCustomer = `INSERT INTO customers (type, login, password, identity) VALUES ($1, $2, $3, $4) RETURNING id, type, status, login, password, identity, created_at, updated_at, deleted`
	updateCustomer = `UPDATE customers SET id=$1, type=$2, status=$3, login=$4, password=$5, identity=$6, created_at=$7, updated_at=$8, deleted=$9 WHERE id = $1 RETURNING id, type, status, login, password, identity, created_at, updated_at, deleted`
	listCustomers  = `SELECT id, type, status, login, password, identity, created_at, updated_at, deleted FROM customers ORDER BY updated_at DESC`
	getCustomer    = `SELECT id, type, status, login, password, identity, created_at, updated_at, deleted FROM customers WHERE id = $1 LIMIT 1`
	deleteCustomer = `DELETE FROM customers WHERE id = $1`
)

func (r *consumerPGRepository) CreateCustomer(ctx context.Context, arg resources.CreateCustomerParams) (resources.Customer, error) {
	row := r.db.QueryRow(ctx, createCustomer, arg.Type, arg.Login, arg.Password, arg.Identity)
	var i resources.Customer
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Status,
		&i.Login,
		&i.Password,
		&i.Identity,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Deleted,
	)
	return i, err
}

func (r *consumerPGRepository) DeleteCustomer(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.Exec(ctx, deleteCustomer, id)
	return err
}

func (r *consumerPGRepository) GetCustomer(ctx context.Context, id uuid.UUID) (resources.Customer, error) {
	row := r.db.QueryRow(ctx, getCustomer, id)
	var i resources.Customer
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Status,
		&i.Login,
		&i.Password,
		&i.Identity,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Deleted,
	)
	return i, err
}

func (r *consumerPGRepository) ListCustomers(ctx context.Context) ([]resources.Customer, error) {
	rows, err := r.db.Query(ctx, listCustomers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []resources.Customer
	for rows.Next() {
		var i resources.Customer
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.Status,
			&i.Login,
			&i.Password,
			&i.Identity,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Deleted,
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

func (r *consumerPGRepository) UpdateCustomer(ctx context.Context, arg resources.UpdateCustomerParams) (resources.Customer, error) {
	row := r.db.QueryRow(ctx, updateCustomer, arg.ID, arg.Type, arg.Status, arg.Login, arg.Password, arg.Identity, arg.CreatedAt, arg.UpdatedAt, arg.Deleted)
	var i resources.Customer
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Status,
		&i.Login,
		&i.Password,
		&i.Identity,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Deleted,
	)
	return i, err
}
