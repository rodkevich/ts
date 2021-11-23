package grpc

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/rodkevich/ts/ticket/internal/models"
	"github.com/rodkevich/ts/ticket/pkg/logger"
	"github.com/rodkevich/ts/ticket/pkg/types"
	"github.com/rodkevich/ts/ticket/proto/ticket/v1"
)

type TicketGrpcService struct {
	v1.UnimplementedTicketServiceServer

	logger    logger.Logger
	useSchema ticket.Invoker
}

func NewTicketGrpcService(logger logger.Logger, useSchema ticket.Invoker) *TicketGrpcService {
	return &TicketGrpcService{logger: logger, useSchema: useSchema}
}

func (s *TicketGrpcService) ListTickets(ctx context.Context, request *v1.ListTicketsRequest) (*v1.ListTicketsResponse, error) {
	c, err := s.useSchema.ListTickets(ctx, nil)
	if err != nil {
		s.logger.Errorf("useSchema.ListTickets: %v", err)
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("%s: %v", "CustomerService.ListCustomers:", err))
	}
	return &v1.ListTicketsResponse{Tickets: c.ToProto()}, nil
}

func (s *TicketGrpcService) CreateTicket(ctx context.Context, r *v1.CreateTicketRequest) (*v1.CreateTicketResponse, error) {
	req := models.Ticket{
		ID:          uuid.MustParse(r.Ticket.GetId()),
		OwnerID:     uuid.MustParse(r.Ticket.OwnerId()),
		NameShort:   "",
		NameExt:     nil,
		Description: nil,
		Amount:      0,
		Price:       0,
		Currency:    0,
		Priority:    types.EnumTicketsPriority(r.Customer.GetType()),
		Published:   false,
		Active:      false,
		CreatedAt:   r.Ticket.GetCreatedAt().AsTime(),
		UpdatedAt:   r.Ticket.GetUpdatedAt().AsTime(),
		Deleted:     false,
	}

	useResp, err := s.useSchema.CreateTicket(ctx, &req)
	if err != nil {
		s.logger.Errorf("useSchema.CreateTicket: %v", err)
		return nil, status.Errorf(codes.AlreadyExists, fmt.Sprintf("%s: %v", "TicketService.CreateTicket:", err))
	}
	resp := v1.CreateTicketResponse{Ticket: useResp.ToProto()}
	return &resp, nil
}
