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

func New(log logger.Logger, ticketRepo ticket.TicketsProprietor) *ticketsController {
	return &ticketsController{ticketPGRepo: ticketRepo, logger: log}
}

func (tc *ticketsController) CreateTicket(ctx context.Context, t *models.Ticket) (*models.Ticket, error) {
	createTicketRepoResp, err := tc.ticketPGRepo.CreateTicket(ctx, t)
	if err != nil {
		tc.logger.Errorf("ticketPGRepo.CreateTicket: %v", err)
		return nil, err
	}

	return createTicketRepoResp, nil
}

func (tc *ticketsController) GetTicket(ctx context.Context, uuid uuid.UUID) (*models.Ticket, error) {
	getTicketRepoResp, err := tc.ticketPGRepo.GetTicket(ctx, uuid)
	if err != nil {
		tc.logger.Errorf("ticketPGRepo.GetTicket: %v", err)
		return nil, err
	}
	return getTicketRepoResp, nil
}

func (tc *ticketsController) ListTickets(ctx context.Context, filter *models.Filter) (*models.TicketsList, error) {
	getTicketsRepoResp, err := tc.ticketPGRepo.ListTickets(ctx, nil)
	if err != nil {
		tc.logger.Errorf("ticketPGRepo.ListTickets: %v", err)
		return nil, err
	}
	return getTicketsRepoResp, nil
}

func (tc *ticketsController) UpdateTicket(ctx context.Context, t *models.Ticket) (*models.Ticket, error) {
	panic("implement me")
}

func (tc *ticketsController) DeleteTicket(ctx context.Context, t *models.Ticket) (*models.Ticket, error) {
	panic("implement me")
}
