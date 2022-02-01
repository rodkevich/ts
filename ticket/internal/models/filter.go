package models

import (
	"errors"

	"github.com/rodkevich/ts/ticket/pkg/filter"
	v1 "github.com/rodkevich/ts/ticket/proto/ticket/v1"
)

const (
	responseItemsDefaultSizePerPage = 10
)

// TicketFilter ...
type TicketFilter struct {
	Base filter.Common

	LastTicketCursor string
}

// func NewTicketFilterFromURL(queries url.Values) *TicketFilter {
// 	f := filter.NewFromURL(queries)
// 	switch {
// 	case queries.Has("LastTicketID"):
// 		fallthrough
// 	case queries.Has("search"):
// 		f.Search = true
// 	}
// 	return &TicketFilter{
// 		Base:                *f,
// 		LastTicketID:        queries.Get("1"),
// 		LastTicketTimestamp: queries.Get("2"),
// 		LastTicketCursor:              queries.Get("3"),
// 	}
// }

// FilterFromRequest ...
func FilterFromRequest(request *v1.ListTicketsRequest) (tf *TicketFilter, err error) {

	f := TicketFilter{
		Base: filter.Common{
			Page:   uint64(request.GetPageSize()),
			Size:   uint64(request.GetPageSize()),
			Search: request.GetSearch(),
			Paging: request.GetPaging(),
		},
	}

	if request.PageSize <= 0 {
		f.Base.Size = responseItemsDefaultSizePerPage
	}

	if request.PageToken != "" {

		println(request.PageToken)
		_, _, err := filter.DecodeCursor(request.PageToken)
		if err != nil {
			return nil, errors.New("invalid-cursor")
		}

		f.LastTicketCursor = request.PageToken

	}

	return &f, nil
}
