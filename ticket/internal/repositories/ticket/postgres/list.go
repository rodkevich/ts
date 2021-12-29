package postgres

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/rodkevich/ts/ticket/internal/models"
)

func (tpg *ticketPG) List(ctx context.Context, id *uuid.UUID, filter *models.TicketFilter) (*models.TicketsList, *uuid.UUID, error) {
	fmt.Printf("pg List: %+v %+v\n", id, filter)

	if !filter.Base.Paging {
		const listTicketsAll = `
		SELECT
		id, owner_id, name_short, name_ext, description, amount, price, currency, priority, published, active, created_at, updated_at, deleted
		FROM tickets
		WHERE tickets.deleted = 'false'
		ORDER BY id DESC
		`
		rows, err := tpg.db.Query(ctx, listTicketsAll)
		if err != nil {
			return nil, nil, err
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
				return nil, nil, err
			}
			rtn = append(rtn, &each)
		}
		if err := rows.Err(); err != nil {
			return nil, nil, err
		}

		return &models.TicketsList{Tickets: rtn}, nil, nil

	} else {
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
			OrderBy("created_at DESC, id DESC")

		if filter != nil {
			if filter.Base.Size > 0 {
				queryBuilder = queryBuilder.Limit(filter.Base.Size)
			}
		}

		if id != nil {
			queryBuilder = queryBuilder.Where(squirrel.LtOrEq{"created_at": "2021-12-27 19:01:05.631728"})
			queryBuilder = queryBuilder.Where(squirrel.Lt{"id": id})
		}

		listTickets, args, err := queryBuilder.ToSql()
		if err != nil {
			return nil, nil, err
		}

		rows, err := tpg.db.Query(
			ctx, listTickets, args...,
		)
		if err != nil {
			return nil, nil, err
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
				return nil, nil, err
			}
			rtn = append(rtn, &each)
		}
		if err := rows.Err(); err != nil {
			return nil, nil, err
		}

		if len(rtn) > 0 {
			lastID := rtn[len(rtn)-1].ID
			fmt.Println(lastID.String())
			return &models.TicketsList{Tickets: rtn}, &lastID, nil
		}

		return &models.TicketsList{Tickets: rtn}, nil, nil

	}

}
