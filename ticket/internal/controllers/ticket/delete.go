/*
 * Copyright 404 1/26/2022.
 *
 *
 */

package ticket

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/rodkevich/ts/ticket/internal/models"
)

func (tc *ticketsController) DeleteTicket(ctx context.Context, id uuid.UUID, hardDelete bool) (*models.Ticket, error) {

	switch hardDelete {
	case true:

		fmt.Printf("controllers DeleteTicket hard-true: %+v %+v\n", id, hardDelete)

		err := tc.ticketPGRepo.Delete(ctx, id, true)
		if err != nil {
			tc.log.Errorf("ticketPGRepo.Delete: %v", err)

			return nil, err
		}

		return nil, nil

	default:

		fmt.Printf("controllers DeleteTicket default: %+v %+v\n", id, hardDelete)

		err := tc.ticketPGRepo.Delete(ctx, id, false)
		if err != nil {
			tc.log.Errorf("ticketPGRepo.Delete: %v", err)

			return nil, err
		}

		return nil, nil

	}
}
