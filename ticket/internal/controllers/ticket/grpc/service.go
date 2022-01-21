package grpc

import (
	"github.com/go-playground/validator/v10"

	"github.com/rodkevich/ts/ticket"
	"github.com/rodkevich/ts/ticket/pkg/logger"

	v1 "github.com/rodkevich/ts/ticket/proto/ticket/v1"
)

type ticketGrpcService struct {
	v1.UnimplementedTicketServiceServer

	log         logger.Logger
	ticketUsage ticket.TicketsUsage
	tagUsage    ticket.TagsUsage
	validate    *validator.Validate
	// ticketTagUsage ticket.TicketTagsController

}

// New ...
func New(logger logger.Logger, useSchema ticket.TicketsUsage, validator *validator.Validate) *ticketGrpcService {

	return &ticketGrpcService{
		log:         logger,
		ticketUsage: useSchema,
		validate:    validator,
	}
}
