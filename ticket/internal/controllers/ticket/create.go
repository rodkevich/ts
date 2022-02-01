/*
 * Copyright 404 1/26/2022.
 *
 *
 */

package ticket

import (
	"context"
	"fmt"

	"github.com/rodkevich/ts/ticket/internal/models"
)

func (tc *ticketsController) CreateTicket(ctx context.Context, t *models.Ticket) (*models.Ticket, error) {
	fmt.Printf("controllers CreateTicket: %+v\n", t)

	createTicketRepoResp, err := tc.ticketPGRepo.Create(ctx, t)
	if err != nil {
		tc.log.Errorf("ticketPGRepo.Create: %v", err)
		return nil, err
	}

	return createTicketRepoResp, nil
}
