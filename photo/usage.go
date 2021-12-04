package photo

import (
	"context"

	"github.com/google/uuid"

	"github.com/rodkevich/ts/photo/internal/models"
)

type PhotosInvoker interface {
	CreatePhoto(context.Context, *models.Photo) (*models.Photo, error)
	GetPhoto(context.Context, uuid.UUID) (*models.Photo, error)
	ListPhotos(context.Context, *models.Photo) (*models.PhotoList, error)
	UpdatePhoto(context.Context, uuid.UUID) (*models.Photo, error)
	DeletePhoto(context.Context, *models.Photo) (*models.Photo, error)
}
