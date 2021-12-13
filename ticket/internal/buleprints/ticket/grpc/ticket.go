package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/rodkevich/ts/ticket"
	"github.com/rodkevich/ts/ticket/internal/models"
	"github.com/rodkevich/ts/ticket/pkg/logger"
	"github.com/rodkevich/ts/ticket/pkg/types"
	"github.com/rodkevich/ts/ticket/proto/ticket/v1"
)

type TicketGrpcService struct {
	v1.UnimplementedTicketServiceServer

	logger      logger.Logger
	ticketUsage ticket.TicketsController
	tagUsage    ticket.TagController
	//ticketTagUsage ticket.TicketTagsController
}

func (tgs TicketGrpcService) CreateTicket(ctx context.Context, request *v1.CreateTicketRequest) (*v1.CreateTicketResponse, error) {
	createTicketUsageResp, err := tgs.ticketUsage.CreateTicket(ctx, &models.Ticket{
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
		tgs.logger.Errorf("ticketUsage.CreateTicket: %v", err)
		return nil, err
	}

	return &v1.CreateTicketResponse{Ticket: createTicketUsageResp.ToProto()}, nil
}

func (tgs TicketGrpcService) ListTickets(ctx context.Context, request *v1.ListTicketsRequest) (*v1.ListTicketsResponse, error) {
	panic("implement me")
}

func (tgs TicketGrpcService) UpdateTicket(ctx context.Context, request *v1.UpdateTicketRequest) (*v1.ListTicketsResponse, error) {
	panic("implement me")
}

func (tgs TicketGrpcService) DeleteTicket(ctx context.Context, request *v1.DeleteTicketRequest) (*v1.DeleteTicketResponse, error) {
	panic("implement me")
}

func New(logger logger.Logger, useSchema ticket.TicketsController) *TicketGrpcService {
	return &TicketGrpcService{logger: logger, ticketUsage: useSchema}
}

// func (s *TicketGrpcService) ListTickets(ctx context.Context, request *ticket_service_v1.ListTicketsRequest) (*ticket_service_v1.ListTicketsResponse, error) {
// 	c, err := s.ticketUsage.ListTickets(ctx, nil)
// 	if err != nil {
// 		s.logger.Errorf("ticketUsage.ListTickets: %v", err)
// 		return nil, status.Errorf(codes.Internal, fmt.Sprintf("%s: %v", "CustomerService.ListCustomers:", err))
// 	}
// 	return &ticket_service_v1.ListTicketsResponse{Tickets: c.ToProto()}, nil
// }
//
// func (s *TicketGrpcService) CreateTicket(ctx context.Context, r *ticket_service_v1.CreateTicketRequest) (*ticket_service_v1.CreateTicketResponse, error) {
// 	req := models.Ticket{
// 		// ID:          uuid.MustParse(r.GetId()),
// 		OwnerID:     uuid.MustParse(r.OwnerId),
// 		NameShort:   "",
// 		NameExt:     nil,
// 		Description: nil,
// 		Amount:      0,
// 		Price:       0,
// 		Currency:    0,
// 		Priority:    types.EnumTicketsPriority(r.GetPriority()),
// 		Published:   false,
// 		Active:      false,
// 		// CreatedAt:   r.Ticket.GetCreatedAt().AsTime(),
// 		// UpdatedAt:   r.Ticket.GetUpdatedAt().AsTime(),
// 		Deleted: false,
// 	}
//
// 	useResp, err := s.ticketUsage.CreateTicket(ctx, &req)
// 	if err != nil {
// 		s.logger.Errorf("ticketUsage.CreateTicket: %v", err)
// 		return nil, status.Errorf(codes.AlreadyExists, fmt.Sprintf("%s: %v", "TicketService.CreateTicket:", err))
// 	}
// 	resp := ticket_service_v1.CreateTicketResponse{Ticket: useResp.ToProto()}
// 	return &resp, nil
// }
