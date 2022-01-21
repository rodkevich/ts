/*
 * Copyright 404 1/21/2022.
 *
 *
 */

package grpc

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/rodkevich/ts/ticket/internal/models"

	"github.com/rodkevich/ts/ticket/pkg/types"
	v1 "github.com/rodkevich/ts/ticket/proto/ticket/v1"
)

// CreateTicket entity with grpc request
func (app ticketGrpcService) CreateTicket(ctx context.Context, request *v1.CreateTicketRequest) (*v1.CreateTicketResponse, error) {
	fmt.Printf("controllers CreateTicket: %+v\n", request)
	parse, err := uuid.Parse(request.GetOwnerId())
	if err != nil {
		return nil, err
	}

	if request != nil {
		// parse grpc struct to known model
		createTicketUsageResp, err := app.ticketUsage.CreateTicket(
			ctx, &models.Ticket{
				OwnerID:     parse,
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
			app.log.Errorf("ticketUsage.Create: %w", err)

			return nil, err
		}

		return &v1.CreateTicketResponse{Ticket: createTicketUsageResp.ToProto()}, nil
	}

	return nil, nil
}
