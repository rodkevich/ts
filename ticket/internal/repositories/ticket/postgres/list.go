package postgres

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/rodkevich/ts/ticket/internal/models"
	"github.com/rodkevich/ts/ticket/pkg/filter"
)

const (
	listTicketsAll = `
		SELECT
		id, owner_id, name_short, name_ext, description, amount, price, currency, priority, published, active, created_at, updated_at, deleted
		FROM tickets
		WHERE tickets.deleted = 'false'
		ORDER BY id DESC
		`

	countTickets = `
		SELECT COUNT(id) as total FROM tickets
		WHERE tickets.deleted = 'false'
		`

	listPaginate = `
		SELECT
		id, owner_id, name_short, name_ext, description, amount, price, currency, priority, published, active, created_at, updated_at, deleted
		FROM tickets
		WHERE (created_at, id) < ($1, $2)
		  AND tickets.deleted = 'false'
		ORDER BY created_at DESC, id DESC
		LIMIT $3;

	`
	listBasic = `
		SELECT
		id, owner_id, name_short, name_ext, description, amount, price, currency, priority, published, active, created_at, updated_at, deleted
		FROM tickets
		WHERE tickets.deleted = 'false'
		ORDER BY created_at DESC, id DESC
		LIMIT $1;
	`
)

// func (tpg *ticketRepositoryPG) List(ctx context.Context, tf models.TicketFilter) (*models.TicketsList, error) {
// 	var totalCounter int
// 	if err := tpg.db.QueryRow(
// 		ctx, countTickets).Scan(&totalCounter); err != nil {
// 		return nil, errors.Wrap(err, "db.Query")
// 	}
//
// 	orderList := []string{"created_at DESC, id DESC"} // should be equal to index
//
// 	// using sq
// 	queryBuilder := sq.Select(
// 		"id", "owner_id", "name_short", "name_ext",
// 		"description", "amount", "price", "currency",
// 		"priority", "published", "active", "created_at",
// 		"updated_at", "deleted",
// 	).
// 		From("tickets").
// 		PlaceholderFormat(sq.Dollar).
// 		Where("tickets.deleted = 'false'").
// 		OrderBy(orderList...)
//
// 	if tf.Base.Size > 0 {
// 		queryBuilder = queryBuilder.Limit(tf.Base.Size)
// 	}
//
// 	if tf.LastTicketCursor != "" {
//
// 		t, u, err := filter.DecodeCursor(tf.LastTicketCursor)
//
// 		if err != nil {
// 			return nil, errors.Wrap(err, "db.List.DecodeCursor")
//
// 		}
// 		queryBuilder = queryBuilder.Where(sq.Expr("(created_at, id ) < ($1, $2)"), t, u)
// 		log.Println(t, u)
//
// 		// queryBuilder = queryBuilder.Where(sq.LtOrEq{"created_at": t})
// 		// queryBuilder = queryBuilder.Where(sq.Lt{"id": u})
//
// 	}
//
// 	// queryBuilder = queryBuilder.Where(sq.LtOrEq{ "created_time": createdCursor })
// 	// queryBuilder = queryBuilder.Where(sq.Lt{ "id": paymentID })
//
// 	listTickets, args, err := queryBuilder.ToSql()
//
// 	log.Println(listTickets, args)
//
// 	if err != nil {
// 		return nil, errors.Wrap(err, "db.queryBuilder.ToSql")
// 	}
//
// 	rows, err := tpg.db.Query(ctx, listTickets, args...)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "db.Query.rows")
// 	}
//
// 	defer rows.Close()
//
// 	var createdTime time.Time
// 	rtn := make([]*models.Ticket, 0)
//
// 	for rows.Next() {
// 		var each models.Ticket
// 		if err := rows.Scan(
// 			&each.ID, &each.OwnerID, &each.NameShort, &each.NameExt,
// 			&each.Description, &each.Amount, &each.Price, &each.Currency,
// 			&each.Priority, &each.Published, &each.Active, &each.CreatedAt,
// 			&each.UpdatedAt, &each.Deleted,
// 		); err != nil {
// 			return nil, err
// 		}
//
// 		createdTime = each.CreatedAt
// 		rtn = append(rtn, &each)
// 	}
//
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
//
// 	if len(rtn) > 0 {
//
// 		cur := filter.EncodeCursor(createdTime, rtn[len(rtn)-1].ID.String())
// 		// lastID := rtn[len(rtn)-1].ID
//
// 		return &models.TicketsList{
// 			Cursor:     cur,
// 			TotalCount: totalCounter,
// 			HasMore:    false,
// 			Tickets:    rtn,
// 		}, nil
// 	}
//
// 	// if no results
// 	return &models.TicketsList{
// 		TotalCount: totalCounter,
// 		HasMore:    false,
// 		Tickets:    nil,
// 	}, nil
// }

func (tpg *ticketRepositoryPG) List(ctx context.Context, tf models.TicketFilter) (tl *models.TicketsList, err error) {
	var counter int
	if err = tpg.db.QueryRow(
		ctx, countTickets,
	).Scan(&counter); err != nil {
		err = errors.Wrap(err, "list.db.query.count")

		return
	}

	rtn := make([]*models.Ticket, 0)
	var createdTime time.Time

	if tf.LastTicketCursor != "" {

		t, u, err := filter.DecodeCursor(tf.LastTicketCursor)

		if err != nil {
			return nil, errors.Wrap(err, "db.List.DecodeCursor")
		}

		rows, err := tpg.db.Query(ctx, listPaginate, t, u, tf.Base.Size)
		if err != nil {
			return nil, errors.Wrap(err, "db.Query.rows")
		}

		defer rows.Close()

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

			createdTime = each.CreatedAt
			rtn = append(rtn, &each)
		}

		if err = rows.Err(); err != nil {
			return nil, err
		}
	}

	tl = new(models.TicketsList)
	if len(rtn) > 0 {
		tl.Cursor = filter.EncodeCursor(createdTime, rtn[len(rtn)-1].ID.String())
		tl.TotalCount = counter
		tl.Tickets = rtn
		return
	}

	return
}
