package controllers

import (
	"context"
	"github.com/google/uuid"
	"github.com/rodkevich/ts/ticket"
	"github.com/rodkevich/ts/ticket/internal/models"
	"github.com/rodkevich/ts/ticket/pkg/logger"
)

type controller struct {
	ticketPGRepo ticket.TicketsProprietor
	logger       logger.Logger
}

func (c controller) CreateTicket(ctx context.Context, m *models.Ticket) (*models.Ticket, error) {
	// TODO implement me
	panic("implement me")
}

func (c controller) GetTicket(ctx context.Context, uuid uuid.UUID) (*models.Ticket, error) {
	// TODO implement me
	panic("implement me")
}

func (c controller) ListTickets(ctx context.Context, m *models.Ticket) (*models.TicketsList, error) {
	// TODO implement me
	panic("implement me")
}

func (c controller) UpdateTicket(ctx context.Context, uuid uuid.UUID) (*models.Ticket, error) {
	// TODO implement me
	panic("implement me")
}

func (c controller) DeleteTicket(ctx context.Context, m *models.Ticket) (*models.Ticket, error) {
	// TODO implement me
	panic("implement me")
}

func NewTicketController(log logger.Logger, ticketRepo ticket.TicketsProprietor) *controller {
	return &controller{ticketPGRepo: ticketRepo, logger: log}
}
