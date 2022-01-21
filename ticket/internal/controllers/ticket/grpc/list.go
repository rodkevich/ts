package grpc

import (
	"context"
	"fmt"

	"github.com/rodkevich/ts/ticket/internal/models"
	v1 "github.com/rodkevich/ts/ticket/proto/ticket/v1"
)

/*
ListTickets handle incoming grpc List request
*/
func (app ticketGrpcService) ListTickets(ctx context.Context, request *v1.ListTicketsRequest) (*v1.ListTicketsResponse, error) {
	/*
		1. достать из запроса условия обработки
		2. собрать фильтр для передачи
		3. вызвать юзадж и получить ответ
		4. проверить на ошибки
		5. забиндить ответ в грписи
	*/
	fmt.Printf("controllers ListTickets: %+v\n", request)

	filters := models.FilterFromRequest(request)

	getTickets, err := app.ticketUsage.ListTickets(ctx, &filters)
	if err != nil {
		app.log.Errorf("ticketUsage.List: %v", err)
		return nil, err
	}

	return &v1.ListTicketsResponse{
		TotalCount:    uint64(getTickets.TotalCount),
		HasMore:       getTickets.HasMore,
		NextPageToken: "",
		Tickets:       getTickets.ToProto(),
	}, nil
}
