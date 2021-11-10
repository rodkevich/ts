package postgres

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type EnumTicketsAdvantagesType string

const (
	EnumTicketsAdvantagesTypeDraft    EnumTicketsAdvantagesType = "Draft"
	EnumTicketsAdvantagesTypeRegular  EnumTicketsAdvantagesType = "Regular"
	EnumTicketsAdvantagesTypePremium  EnumTicketsAdvantagesType = "Premium"
	EnumTicketsAdvantagesTypePromoted EnumTicketsAdvantagesType = "Promoted"
)

func (e *EnumTicketsAdvantagesType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = EnumTicketsAdvantagesType(s)
	case string:
		*e = EnumTicketsAdvantagesType(s)
	default:
		return fmt.Errorf("unsupported scan type for EnumTicketsAdvantagesType: %T", src)
	}
	return nil
}

type Ticket struct {
	ID          uuid.UUID                 `json:"id,omitempty"`
	OwnerID     uuid.UUID                 `json:"owner_id,omitempty"`
	NameShort   string                    `json:"name_short,omitempty"`
	NameExt     string                    `json:"name_ext,omitempty"`
	Description *string                   `json:"description,omitempty"`
	Amount      int32                     `json:"amount,omitempty"`
	Price       float64                   `json:"price,omitempty"`
	Currency    int32                     `json:"currency,omitempty"`
	Active      bool                      `json:"active,omitempty"`
	Advantage   EnumTicketsAdvantagesType `json:"advantage,omitempty"`
	PublishedAt *time.Time                `json:"published_at,omitempty"`
	CreatedAt   time.Time                 `json:"created_at"`
	UpdatedAt   time.Time                 `json:"updated_at"`
	DeletedAt   *time.Time                `json:"deleted_at,omitempty"`
}
