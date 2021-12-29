package grpc

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/go-playground/validator/v10"

	"github.com/rodkevich/ts/ticket"
	"github.com/rodkevich/ts/ticket/internal/models"
	"github.com/rodkevich/ts/ticket/pkg/filter"
	"github.com/rodkevich/ts/ticket/pkg/logger"
	"github.com/rodkevich/ts/ticket/pkg/types"
	"github.com/rodkevich/ts/ticket/proto/ticket/v1"
)

type ticketGrpcService struct {
	v1.UnimplementedTicketServiceServer

	log         logger.Logger
	ticketUsage ticket.TicketsController
	tagUsage    ticket.TagController
	// ticketTagUsage ticket.TicketTagsController
	validate *validator.Validate
}

func New(logger logger.Logger, useSchema ticket.TicketsController, validator *validator.Validate) *ticketGrpcService {

	return &ticketGrpcService{
		log:         logger,
		ticketUsage: useSchema,
		validate:    validator,
	}
}

func (app ticketGrpcService) CreateTicket(ctx context.Context, request *v1.CreateTicketRequest) (*v1.CreateTicketResponse, error) {
	fmt.Printf("handlers CreateTicket: %+v\n", request)

	if request != nil {
		// parse grpc struct to known model
		createTicketUsageResp, err := app.ticketUsage.CreateTicket(
			ctx, &models.Ticket{
				OwnerID:     uuid.MustParse(request.GetOwnerId()),
				NameShort:   request.GetNameShort(),
				NameExt:     &request.NameExt,
				Description: &request.Description,
				Amount:      request.GetAmount(),
				Price:       request.GetPrice(),
				Currency:    request.GetCurrency(),
				Priority:    types.EnumTicketsPriority(request.GetPriority()),
				Published:   request.GetPublished(),
			})
		if err != nil {
			app.log.Errorf("ticketUsage.Create: %v", err)

			return nil, err
		}

		return &v1.CreateTicketResponse{Ticket: createTicketUsageResp.ToProto()}, nil
	}

	return nil, nil
}

func (app ticketGrpcService) ListTickets(ctx context.Context, request *v1.ListTicketsRequest) (*v1.ListTicketsResponse, error) {
	fmt.Printf("handlers ListTickets: %+v\n", request)

	filters := models.TicketFilter{
		Base: filter.Common{
			Extended: request.GetExtended(),
			Search:   request.GetSearch(),
			Page:     uint64(request.GetPageSize()),
			Size:     uint64(request.PageSize),
			Paging:   request.GetPaging(),
		},
		LastId: request.GetId(),
		Field2: "",
		Field3: "",
	}

	getTickets, err := app.ticketUsage.ListTickets(ctx, &filters)
	if err != nil {
		app.log.Errorf("ticketUsage.List: %v", err)
		return nil, err
	}

	var rtn = make([]*v1.Ticket, 0, len(getTickets.Tickets))
	for _, x := range getTickets.Tickets {
		rtn = append(rtn, x.ToProto())
	}

	return &v1.ListTicketsResponse{
		Tickets:       rtn,
		NextPageToken: "",
	}, nil
}

func (app ticketGrpcService) GetTicket(ctx context.Context, r *v1.GetTicketRequest) (*v1.ListTicketsResponse, error) {
	fmt.Printf("handlers GetTicket: %+v\n", r)

	getTicket, err := app.ticketUsage.GetTicket(ctx, uuid.MustParse(r.GetId()))
	if err != nil {
		app.log.Errorf("ticketUsage.Get: %v", err)
		return nil, err
	}

	var rtn = make([]*v1.Ticket, 0, 1)

	return &v1.ListTicketsResponse{
		Tickets:       append(rtn, getTicket.ToProto()),
		NextPageToken: "",
	}, nil
}

func (app ticketGrpcService) UpdateTicket(ctx context.Context, request *v1.UpdateTicketRequest) (*v1.ListTicketsResponse, error) {
	fmt.Printf("handlers UpdateTicket: %+v\n", request)
	panic("implement me")
}

func (app ticketGrpcService) DeleteTicket(ctx context.Context, req *v1.DeleteTicketRequest) (*v1.DeleteTicketResponse, error) {
	fmt.Printf("handlers DeleteTicket: %+v\n", req)

	id := uuid.MustParse(req.GetId())
	isHardDelete := req.GetHard()

	_, err := app.ticketUsage.DeleteTicket(ctx, id, isHardDelete)
	if err != nil {
		app.log.Errorf("ticketUsage.Delete: %v", err)
		return nil, err
	}

	return &v1.DeleteTicketResponse{}, nil
}
