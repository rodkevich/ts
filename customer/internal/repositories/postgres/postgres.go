package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/rodkevich/ts/customer/internal/models"
)

type customerPG struct {
	db *pgxpool.Pool
}

func NewCustomer(db *pgxpool.Pool) *customerPG {
	return &customerPG{db: db}
}

const (
	createCustomer = `INSERT INTO customers (type, login, password, identity) VALUES ($1, $2, $3, $4) RETURNING id, type, status, login, password, identity, created_at, updated_at, deleted`
	updateCustomer = `UPDATE customers SET id=$1, type=$2, status=$3, login=$4, password=$5, identity=$6, created_at=$7, updated_at=$8, deleted=$9 WHERE id = $1 RETURNING id, type, status, login, password, identity, created_at, updated_at, deleted`
	listCustomers  = `SELECT id, type, status, login, password, identity, created_at, updated_at, deleted FROM customers ORDER BY updated_at DESC`
	getCustomer    = `SELECT id, type, status, login, password, identity, created_at, updated_at, deleted FROM customers WHERE id = $1 LIMIT 1`
	deleteCustomer = `DELETE FROM customers WHERE id = $1`
)

func (r *customerPG) CreateCustomer(ctx context.Context, arg models.CreateCustomerParams) (*models.Customer, error) {
	row := r.db.QueryRow(
		ctx, createCustomer, arg.Type, arg.Login, arg.Password, arg.Identity,
	)

	var rtn models.Customer
	err := row.Scan(
		&rtn.ID, &rtn.Type, &rtn.Status, &rtn.Login,
		&rtn.Password, &rtn.Identity, &rtn.CreatedAt,
		&rtn.UpdatedAt, &rtn.Deleted,
	)
	return &rtn, err
}

func (r *customerPG) GetCustomer(ctx context.Context, id uuid.UUID) (*models.Customer, error) {
	row := r.db.QueryRow(
		ctx, getCustomer, id,
	)

	var rtn models.Customer
	err := row.Scan(
		&rtn.ID, &rtn.Type, &rtn.Status,
		&rtn.Login, &rtn.Password, &rtn.Identity,
		&rtn.CreatedAt, &rtn.UpdatedAt, &rtn.Deleted,
	)
	return &rtn, err
}

func (r *customerPG) DeleteCustomer(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.Exec(
		ctx, deleteCustomer, id,
	)
	return err
}

func (r *customerPG) ListCustomers(ctx context.Context) (*models.CustomersList, error) {
	rows, err := r.db.Query(
		ctx, listCustomers,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rtn := make([]*models.Customer, 0)
	for rows.Next() {
		var each models.Customer
		if err := rows.Scan(
			&each.ID,
			&each.Type,
			&each.Status,
			&each.Login,
			&each.Password,
			&each.Identity,
			&each.CreatedAt,
			&each.UpdatedAt,
			&each.Deleted,
		); err != nil {
			return nil, err
		}
		rtn = append(rtn, &each)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &models.CustomersList{Customers: rtn}, nil
}

func (r *customerPG) UpdateCustomer(ctx context.Context, arg models.UpdateCustomerParams) (*models.Customer, error) {
	row := r.db.QueryRow(
		ctx, updateCustomer, arg.ID, arg.Type, arg.Status, arg.Login,
		arg.Password, arg.Identity, arg.CreatedAt, arg.UpdatedAt, arg.Deleted,
	)
	var rtn models.Customer
	err := row.Scan(
		&rtn.ID, &rtn.Type, &rtn.Status, &rtn.Login, &rtn.Password,
		&rtn.Identity, &rtn.CreatedAt, &rtn.UpdatedAt, &rtn.Deleted,
	)
	return &rtn, err
}
