package grpc

import (
	"context"
	"fmt"

	"github.com/rodkevich/ts/photo"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/rodkevich/ts/photo/internal/models"
	"github.com/rodkevich/ts/photo/pkg/logger"
	"github.com/rodkevich/ts/photo/proto/photo/v1"
)

type PhotoGrpcService struct {
	v1.UnimplementedPhotoServiceServer

	logger     logger.Logger
	photoUsage photo.PhotosInvoker
}

func NewPhotoGrpcService(logger logger.Logger, useSchema photo.PhotosInvoker) *PhotoGrpcService {
	return &PhotoGrpcService{logger: logger, photoUsage: useSchema}
}

func (s *PhotoGrpcService) ListPhotos(ctx context.Context, request *v1.ListPhotosRequest) (*v1.ListPhotosResponse, error) {
	c, err := s.photoUsage.ListPhotos(ctx, nil)
	if err != nil {
		s.logger.Errorf("photoUsage.ListPhotos: %v", err)
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("%s: %v", "PhotosService.ListPhotos:", err))
	}
	return &v1.ListPhotosResponse{Photos: c.ToProto()}, nil
}

func (s *PhotoGrpcService) CreateTicket(ctx context.Context, r *v1.CreatePhotoRequest) (*v1.CreatePhotoResponse, error) {
	req := models.Photo{
		// OwnerID:     uuid.MustParse(r.OwnerId()),
		// Priority:    types.EnumTicketsPriority(r.Customer.GetType()),
		// CreatedAt:   r.Ticket.GetCreatedAt().AsTime(),
		// UpdatedAt:   r.Ticket.GetUpdatedAt().AsTime(),
	}

	useResp, err := s.photoUsage.CreatePhoto(ctx, &req)
	if err != nil {
		s.logger.Errorf("ticketUsage.CreateTicket: %v", err)
		return nil, status.Errorf(codes.AlreadyExists, fmt.Sprintf("%s: %v", "TicketService.CreateTicket:", err))
	}
	resp := v1.CreatePhotoResponse{Photo: useResp.ToProto()}
	return &resp, nil
}
