package photo

import (
	"context"

	"github.com/google/uuid"

	"github.com/rodkevich/ts/photo/internal/models"
)

type PhotosProprietor interface {
	CreatePhoto(ctx context.Context, arg models.CreatePhotoParams) (*models.Photo, error)
	GetPhoto(ctx context.Context, id uuid.UUID) (*models.Photo, error)
	ListPhotos(ctx context.Context) (*models.PhotoList, error)
	UpdatePhoto(ctx context.Context, arg models.UpdatePhotoParams, id uuid.UUID) (*models.Photo, error)

	MarkPhotoAsDeleted(ctx context.Context, deleted bool, id uuid.UUID) (*models.Photo, error)
	DeletePhoto(ctx context.Context, id uuid.UUID) error
}
