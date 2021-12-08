package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/rodkevich/ts/ticket"
	"github.com/rodkevich/ts/ticket/internal/models"
	"github.com/rodkevich/ts/ticket/pkg/logger"
	ticket_service_v1 "github.com/rodkevich/ts/ticket/proto/ticket/v1"
)

type TicketGrpcService struct {
	ticket_service_v1.UnimplementedTicketServiceServer

	logger      logger.Logger
	ticketUsage ticket.TicketsInvoker
	tagUsage    ticket.TagsInvoker
}

func (t TicketGrpcService) CreateTag(ctx context.Context, tag *models.Tag) (*models.Tag, error) {
	// TODO implement me
	panic("implement me")
}

func (t TicketGrpcService) GetTag(ctx context.Context, uuid uuid.UUID) (*models.Tag, error) {
	// TODO implement me
	panic("implement me")
}

func (t TicketGrpcService) ListTags(ctx context.Context, tag *models.Tag) (*models.TagList, error) {
	// TODO implement me
	panic("implement me")
}

func (t TicketGrpcService) UpdateTag(ctx context.Context, uuid uuid.UUID) (*models.Tag, error) {
	// TODO implement me
	panic("implement me")
}

func (t TicketGrpcService) DeleteTag(ctx context.Context, tag *models.Tag) (*models.Tag, error) {
	// TODO implement me
	panic("implement me")
}

func NewTicketGrpcService(logger logger.Logger, useSchema ticket.TicketsInvoker) *TicketGrpcService {
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
