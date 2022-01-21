package models

import (
	"net/url"

	"github.com/rodkevich/ts/ticket/pkg/filter"
	v1 "github.com/rodkevich/ts/ticket/proto/ticket/v1"
)

type TicketFilter struct {
	Base filter.Common

	LastId string
	Field2 string
	Field3 string
}

func NewTicketFilterFromURL(queries url.Values) *TicketFilter {
	f := filter.NewFromURL(queries)
	switch {
	case queries.Has("LastId"):
		fallthrough
	case queries.Has("search"):
		f.Search = true
	}
	return &TicketFilter{
		Base:   *f,
		LastId: queries.Get("1"),
		Field2: queries.Get("2"),
		Field3: queries.Get("3"),
	}
}

func FilterFromRequest(request *v1.ListTicketsRequest) TicketFilter {
	rtn := TicketFilter{
		Base: filter.Common{
			Page:   uint64(request.GetPageSize()),
			Size:   uint64(request.GetPageSize()),
			Search: request.GetSearch(),
			Paging: request.GetPaging(),
		},
		LastId: request.GetId(),
		Field2: "",
		Field3: "",
	}
	return rtn
}
