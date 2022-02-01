/*
 * Copyright 404 1/26/2022.
 *
 *
 */

package ticket

import (
	"context"

	"github.com/rodkevich/ts/ticket/internal/models"
)

func (tc *ticketsController) ListTickets(ctx context.Context, tf *models.TicketFilter) (*models.TicketsList, error) {

	// business logic dispatcher search/list/list_paginated
	switch tf.Base.Search {

	case true:

		repoSearch, _, err := tc.ticketPGRepo.Search(ctx, nil, tf)
		if err != nil {
			tc.log.Errorf("ticketPGRepo.Search: %v", err)
			return nil, err
		}

		return repoSearch, nil

	default:

		if tf.LastTicketCursor != "" {
			repoList, err := tc.ticketPGRepo.List(ctx, *tf)
			if err != nil {
				tc.log.Errorf("ticketPGRepo.List: %v", err)

				return nil, err
			}

			return repoList, nil
		}

		repoList, err := tc.ticketPGRepo.List(ctx, *tf)
		if err != nil {
			tc.log.Errorf("ticketPGRepo.List: %v", err)

			return nil, err
		}

		return repoList, nil
	}
}
