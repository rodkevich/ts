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

	v1 "github.com/rodkevich/ts/ticket/proto/ticket/v1"
)

func (app ticketGrpcService) GetTicket(ctx context.Context, r *v1.GetTicketRequest) (*v1.ListTicketsResponse, error) {
	fmt.Printf("controllers GetTicket: %+v\n", r)

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
