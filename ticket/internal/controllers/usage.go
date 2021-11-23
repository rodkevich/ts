package controllers

import (
	"github.com/rodkevich/ts/ticket"
	"github.com/rodkevich/ts/ticket/pkg/logger"
)

type controller struct {
	ticketPGRepo ticket.ProprietorTickets
	logger       logger.Logger
}

func NewTicketController(log logger.Logger, ticketRepo ticket.ProprietorTickets) *controller {
	return &controller{ticketPGRepo: ticketRepo, logger: log}
}
