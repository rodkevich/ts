package handlers

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/rodkevich/ts/ticket"
	"github.com/rodkevich/ts/ticket/internal/models"
	"github.com/rodkevich/ts/ticket/pkg/logger"
)

type ticketsUsage struct {
	log          logger.Logger
	ticketPGRepo ticket.TicketsProprietor
}

func New(log logger.Logger, ticketRepo ticket.TicketsProprietor) *ticketsUsage {
	return &ticketsUsage{ticketPGRepo: ticketRepo, log: log}
}

func (app *ticketsUsage) CreateTicket(ctx context.Context, t *models.Ticket) (*models.Ticket, error) {
	fmt.Printf("handlers CreateTicket: %+v\n", t)

	createTicketRepoResp, err := app.ticketPGRepo.Create(ctx, t)
	if err != nil {
		app.log.Errorf("ticketPGRepo.Create: %v", err)
		return nil, err
	}

	return createTicketRepoResp, nil
}

func (app *ticketsUsage) GetTicket(ctx context.Context, uuid uuid.UUID) (*models.Ticket, error) {
	getTicketRepoResp, err := app.ticketPGRepo.Get(ctx, uuid)
	if err != nil {
		app.log.Errorf("ticketPGRepo.Get: %v", err)
		return nil, err
	}
	return getTicketRepoResp, nil
}

func (app *ticketsUsage) ListTickets(ctx context.Context, filter *models.TicketFilter) (*models.TicketsList, error) {

	lastID := uuid.MustParse(filter.LastId)

	switch filter.Base.Search {
	case true:
		repoSearch, _, err := app.ticketPGRepo.Search(ctx, nil, filter)
		if err != nil {
			app.log.Errorf("ticketPGRepo.Search: %v", err)
			return nil, err
		}
		return repoSearch, nil

	default:
		repoList, _, err := app.ticketPGRepo.List(ctx, &lastID, filter)
		if err != nil {
			app.log.Errorf("ticketPGRepo.List: %v", err)
			return nil, err
		}

		return repoList, nil
	}
}

func (app *ticketsUsage) UpdateTicket(ctx context.Context, t *models.Ticket) (*models.Ticket, error) {
	panic("implement me")
}

func (app *ticketsUsage) DeleteTicket(ctx context.Context, id uuid.UUID, hardDelete bool) (*models.Ticket, error) {

	switch hardDelete {
	case true:
		fmt.Printf("handlers DeleteTicket hard-true: %+v %+v\n", id, hardDelete)

		err := app.ticketPGRepo.Delete(ctx, id, true)
		if err != nil {
			app.log.Errorf("ticketPGRepo.Delete: %v", err)
			return nil, err
		}

		return nil, nil

	default:
		fmt.Printf("handlers DeleteTicket default: %+v %+v\n", id, hardDelete)

		err := app.ticketPGRepo.Delete(ctx, id, false)
		if err != nil {
			app.log.Errorf("ticketPGRepo.Delete: %v", err)
			return nil, err
		}

		return nil, nil

	}
}
