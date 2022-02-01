/*
 * Copyright 404 1/26/2022.
 *
 *
 */

package ticket

import (
	"context"

	"github.com/google/uuid"

	"github.com/rodkevich/ts/ticket/internal/models"
)

func (tc *ticketsController) GetTicket(ctx context.Context, uuid uuid.UUID) (*models.Ticket, error) {
	getTicketRepoResp, err := tc.ticketPGRepo.Get(ctx, uuid)
	if err != nil {
		tc.log.Errorf("ticketPGRepo.Get: %v", err)

		return nil, err
	}

	return getTicketRepoResp, nil
}
