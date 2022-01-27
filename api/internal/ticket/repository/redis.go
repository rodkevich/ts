package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/rodkevich/ts/api/internal/models"
)

const (
	prefix     = "tickets:"
	expiration = time.Second * 3600
)

type ticketRedisRepo struct {
	redisConn *redis.Client
}

func NewTicketRedisRepo(redisConn *redis.Client) TicketRedisRepo {
	return &ticketRedisRepo{redisConn: redisConn}
}

func (trr *ticketRedisRepo) GetTicketByID(ctx context.Context, id uuid.UUID) (*models.Ticket, error) {

	result, err := trr.redisConn.Get(ctx, trr.createKey(id)).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "ticketRedisRepo.GetTicketByID")
	}

	var res models.Ticket
	if err := json.Unmarshal(result, &res); err != nil {
		return nil, errors.Wrap(err, "ticketRedisRepo.GetTicketByID.json.Unmarshal")
	}

	return &res, nil
}

func (trr *ticketRedisRepo) SetTicket(ctx context.Context, t *models.Ticket) error {

	ticketBytes, err := json.Marshal(t)
	if err != nil {
		return errors.Wrap(err, "ticketRedisRepo.SetTicket.Marshal")
	}

	if err := trr.redisConn.SetEX(ctx, trr.createKey(t.ID), string(ticketBytes), expiration).Err(); err != nil {
		return errors.Wrap(err, "ticketRedisRepo.SetTicket.SetEX")
	}

	return nil
}

func (trr *ticketRedisRepo) DeleteTicket(ctx context.Context, id uuid.UUID) error {

	if err := trr.redisConn.Del(ctx, trr.createKey(id)).Err(); err != nil {
		return errors.Wrap(err, "ticketRedisRepo.DeleteTicket.Del")
	}

	return nil
}

func (trr *ticketRedisRepo) createKey(hotelID uuid.UUID) string {
	return fmt.Sprintf("%s: %s", prefix, hotelID.String())
}
