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

func (app ticketGrpcService) DeleteTicket(ctx context.Context, req *v1.DeleteTicketRequest) (*v1.DeleteTicketResponse, error) {
	fmt.Printf("controllers DeleteTicket: %+v\n", req)

	id := uuid.MustParse(req.GetId())
	isHardDelete := req.GetHard()

	_, err := app.ticketUsage.DeleteTicket(ctx, id, isHardDelete)
	if err != nil {
		app.log.Errorf("ticketUsage.Delete: %v", err)
		return nil, err
	}

	return &v1.DeleteTicketResponse{}, nil
}
