package controllers

import (
	"context"
	"github.com/google/uuid"

	"github.com/rodkevich/ts/ticket"
	"github.com/rodkevich/ts/ticket/internal/models"
	"github.com/rodkevich/ts/ticket/pkg/logger"
)

type ticketsController struct {
	logger       logger.Logger
	ticketPGRepo ticket.TicketsProprietor
}

func (tc *ticketsController) CreateTicket(ctx context.Context, m *models.Ticket) (*models.Ticket, error) {
	createTicketDBResp, err := tc.ticketPGRepo.CreateTicket(ctx, m)
	if err != nil {
		tc.logger.Errorf("ticketPGRepo.CreateTicket: %v", err)
		return nil, err
	}
	return createTicketDBResp, nil
}

func (tc *ticketsController) GetTicket(ctx context.Context, uuid uuid.UUID) (*models.Ticket, error) {
	panic("implement me")
}

func (tc *ticketsController) ListTickets(ctx context.Context, filter *models.Filter) (*models.TicketsList, error) {
	panic("implement me")
}

func (tc *ticketsController) UpdateTicket(ctx context.Context, uuid uuid.UUID) (*models.Ticket, error) {
	panic("implement me")
}

func (tc *ticketsController) DeleteTicket(ctx context.Context, m *models.Ticket) (*models.Ticket, error) {
	panic("implement me")
}

func New(log logger.Logger, ticketRepo ticket.TicketsProprietor) *ticketsController {
	return &ticketsController{ticketPGRepo: ticketRepo, logger: log}
}
