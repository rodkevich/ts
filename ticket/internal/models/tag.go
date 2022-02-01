package models

import (
	"time"

	"github.com/google/uuid"

	"github.com/rodkevich/ts/ticket/pkg/types"
	v1 "github.com/rodkevich/ts/ticket/proto/tag/v1"
)

type (
	// Tag ...
	Tag struct {
		ID          uuid.UUID
		Name        string
		Description *string
		CreatedAt   time.Time
		UpdatedAt   time.Time
		Deleted     bool
	}

	// TagList ...
	TagList struct {
		Tags []*Tag `json:"tags"`
	}

	// CreateTagParams ...
	CreateTagParams struct {
		Name        string
		Description types.NullString
	}

	// UpdateTagParams ...
	UpdateTagParams struct {
		ID          uuid.UUID
		Name        string
		Description *string
		Deleted     bool
	}
)

// ToProto ...
func (t *Tag) ToProto() *v1.Tag {
	return &v1.Tag{
		Id:          "",
		Name:        "",
		Description: "",
		CreatedAt:   nil,
		UpdatedAt:   nil,
		Deleted:     false,
	}
}

// ToProto ...
func (tl *TagList) ToProto() []*v1.Tag {
	tagsList := make([]*v1.Tag, 0, len(tl.Tags))
	for _, tag := range tl.Tags {
		tagsList = append(tagsList, tag.ToProto())
	}
	return tagsList
}
