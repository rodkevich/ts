package models

import (
	"time"

	"github.com/google/uuid"

	"github.com/rodkevich/ts/photo/pkg/types"
	"github.com/rodkevich/ts/photo/proto/photo/v1"
)

type Photo struct {
	ID          uuid.UUID
	Type        types.EnumPhotosType
	SizeKb      float64
	UploadName  types.NullString
	ImageUrl    string
	Description types.NullString
	OwnerID     uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Deleted     bool
}

func (p *Photo) ToProto() *v1.Photo {
	return &v1.Photo{
		Id:          "",
		Type:        "",
		SizeKb:      0,
		UploadName:  "",
		Description: "",
		OwnerId:     "",
		CreatedAt:   nil,
		UpdatedAt:   nil,
		Deleted:     false,
	}
}

type PhotoList struct {
	Photos []*Photo `json:"photos"`
}

func (pl *PhotoList) ToProto() []*v1.Photo {
	photosList := make([]*v1.Photo, 0, len(pl.Photos))
	for _, photo := range pl.Photos {
		photosList = append(photosList, photo.ToProto())
	}
	return photosList
}

type CreatePhotoParams struct {
	Type        types.EnumPhotosType
	SizeKb      float64
	UploadName  types.NullString
	Description types.NullString
	ImageUrl    string
	OwnerID     uuid.UUID
}

type UpdatePhotoParams struct {
	Type        types.EnumPhotosType
	SizeKb      float64
	UploadName  types.NullString
	ImageUrl    string
	Description types.NullString
	OwnerID     uuid.UUID
}
