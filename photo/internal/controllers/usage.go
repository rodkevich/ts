package controllers

import (
	"context"
	"github.com/google/uuid"
	"github.com/rodkevich/ts/photo"
	"github.com/rodkevich/ts/photo/internal/models"
	"github.com/rodkevich/ts/photo/pkg/logger"
)

type controller struct {
	photoPGRepo photo.PhotosProprietor
	logger      logger.Logger
}

func NewPhotoController(log logger.Logger, photoRepo photo.PhotosProprietor) *controller {
	return &controller{photoPGRepo: photoRepo, logger: log}
}

func (c controller) CreatePhoto(ctx context.Context, m *models.Photo) (*models.Photo, error) {
	// TODO implement me
	panic("implement me")
}

func (c controller) GetPhoto(ctx context.Context, uuid uuid.UUID) (*models.Photo, error) {
	// TODO implement me
	panic("implement me")
}

func (c controller) ListPhotos(ctx context.Context, m *models.Photo) (*models.PhotoList, error) {
	// TODO implement me
	panic("implement me")
}

func (c controller) UpdatePhoto(ctx context.Context, uuid uuid.UUID) (*models.Photo, error) {
	// TODO implement me
	panic("implement me")
}

func (c controller) DeletePhoto(ctx context.Context, m *models.Photo) (*models.Photo, error) {
	// TODO implement me
	panic("implement me")
}
