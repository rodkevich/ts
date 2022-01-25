/*
 * Copyright 404 1/24/2022.
 *
 *
 */

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
	ticketUsage ticket.TicketsController
	tagUsage    ticket.TagsController
	validate    *validator.Validate
	// ticketTagUsage ticket.TicketTagsController

}

// New Grpc service for ticket
func New(logger logger.Logger, useSchema ticket.TicketsController, validator *validator.Validate) v1.TicketServiceServer {

	return &ticketGrpcService{
		log:         logger,
		ticketUsage: useSchema,
		validate:    validator,
	}

}
