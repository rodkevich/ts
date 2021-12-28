package models

import (
	"net/url"

	"github.com/rodkevich/ts/ticket/pkg/filter"
)

type TicketFilter struct {
	Base filter.Common

	LastId string
	Field2 string
	Field3 string
}

func NewTicketFilterFromURL(queries url.Values) *TicketFilter {
	f := filter.NewFromURL(queries)
	return &TicketFilter{
		Base:   *f,
		LastId: queries.Get("1"),
		Field2: queries.Get("2"),
		Field3: queries.Get("3"),
	}
}
