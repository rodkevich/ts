/*
 * Copyright 404 1/21/2022.
 *
 *
 */

package grpc

import (
	"context"
	"fmt"

	v1 "github.com/rodkevich/ts/ticket/proto/ticket/v1"
)

func (app ticketGrpcService) UpdateTicket(ctx context.Context, request *v1.UpdateTicketRequest) (*v1.ListTicketsResponse, error) {
	fmt.Printf("controllers UpdateTicket: %+v\n", request)
	panic("implement me")
}
