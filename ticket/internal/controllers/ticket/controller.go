/*
 * Copyright 404 1/26/2022.
 *
 *
 */

package ticket

import (
	"github.com/rodkevich/ts/ticket"
	"github.com/rodkevich/ts/ticket/pkg/logger"
)

type ticketsController struct {
	log          logger.Logger
	ticketPGRepo ticket.TicketsRepositoryIFace
}

// New usage controller
func New(log logger.Logger, ticketRepo ticket.TicketsRepositoryIFace) ticket.TicketsController {
	return &ticketsController{log, ticketRepo}
}
