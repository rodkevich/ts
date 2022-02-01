/*
 * Copyright 404 1/27/2022.
 *
 *
 */

package ticket

import (
	"context"

	"github.com/rodkevich/ts/ticket/internal/models"
)

func (tc *ticketsController) SearchTickets(ctx context.Context, tf *models.TicketFilter) (*models.TicketsList, error) {
	repoSearch, _, err := tc.ticketPGRepo.Search(ctx, nil, tf)

	if err != nil {
		tc.log.Errorf("ticketPGRepo.Search: %v", err)
		return nil, err
	}

	return repoSearch, nil
}
