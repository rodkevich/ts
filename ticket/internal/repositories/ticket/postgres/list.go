package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
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
		SELECT COUNT(*) as total FROM tickets
		WHERE tickets.deleted = 'false'
		`
)

func (tpg *ticketRepositoryPG) List(ctx context.Context, ticketFilter models.TicketFilter) (*models.TicketsList, *uuid.UUID, error) {
	fmt.Printf("pg List request: %+v\n", ticketFilter)

	var totalCounter int
	if err := tpg.db.QueryRow(ctx, countTickets).Scan(&totalCounter); err != nil {

		return nil, nil, errors.Wrap(err, "db.Query")
	}

	orderList := []string{"created_at DESC, id DESC"}

	// using squirrel
	queryBuilder := squirrel.Select(
		"id", "owner_id", "name_short", "name_ext",
		"description", "amount", "price", "currency",
		"priority", "published", "active", "created_at",
		"updated_at", "deleted",
	).
		From("tickets").
		PlaceholderFormat(squirrel.Dollar).
		Where("tickets.deleted = 'false'").
		OrderBy(orderList...)

	if ticketFilter.Base.Size > 0 {
		queryBuilder = queryBuilder.Limit(ticketFilter.Base.Size)
	}

	if ticketFilter.LastTicketTimestamp != "" {
		println(ticketFilter.LastTicketTimestamp)
		queryBuilder = queryBuilder.Where(squirrel.LtOrEq{"created_at": ticketFilter.LastTicketTimestamp})
	}

	if ticketFilter.LastTicketID != "" {
		queryBuilder = queryBuilder.Where(squirrel.Lt{"id": ticketFilter.LastTicketID})
	}

	// queryBuilder = queryBuilder.Where(sq.LtOrEq{ "created_time": createdCursor })
	// queryBuilder = queryBuilder.Where(sq.Lt{ "id": paymentID })

	listTickets, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, nil, errors.Wrap(err, "db.queryBuilder.ToSql")
	}

	rows, err := tpg.db.Query(ctx, listTickets, args...)
	if err != nil {
		return nil, nil, errors.Wrap(err, "db.Query.rows")
	}

	defer rows.Close()

	var createdTime time.Time
	rtn := make([]*models.Ticket, 0)

	for rows.Next() {
		var each models.Ticket
		if err := rows.Scan(
			&each.ID, &each.OwnerID, &each.NameShort, &each.NameExt,
			&each.Description, &each.Amount, &each.Price, &each.Currency,
			&each.Priority, &each.Published, &each.Active, &each.CreatedAt,
			&each.UpdatedAt, &each.Deleted,
		); err != nil {
			return nil, nil, err
		}

		createdTime = each.CreatedAt
		rtn = append(rtn, &each)
	}

	if err := rows.Err(); err != nil {
		return nil, nil, err
	}

	if len(rtn) > 0 {

		cur := filter.EncodeCursor(createdTime, rtn[len(rtn)-1].ID.String())
		lastID := rtn[len(rtn)-1].ID
		log.Println("-- ID --", lastID)
		log.Println("-- cur --", cur)
		log.Println("-- createdTime --", createdTime)

		return &models.TicketsList{
			Cursor:     cur,
			TotalCount: totalCounter,
			HasMore:    false,
			Tickets:    rtn,
		}, &lastID, nil
	}

	// if no results
	return &models.TicketsList{
		TotalCount: totalCounter,
		HasMore:    false,
		Tickets:    nil,
	}, nil, nil
}
