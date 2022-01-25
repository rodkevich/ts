package models

import (
	"log"

	"github.com/rodkevich/ts/ticket/pkg/filter"
	v1 "github.com/rodkevich/ts/ticket/proto/ticket/v1"
)

type TicketFilter struct {
	Base filter.Common

	LastTicketID        string
	LastTicketTimestamp string
	LastTicketCursor    string
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

func FilterFromRequest(request *v1.ListTicketsRequest) TicketFilter {
	rtn := TicketFilter{
		Base: filter.Common{
			Page:   uint64(request.GetPageSize()),
			Size:   uint64(request.GetPageSize()),
			Search: request.GetSearch(),
			Paging: request.GetPaging(),
		},
		LastTicketID:        "",
		LastTicketTimestamp: "",
		LastTicketCursor:    "",
	}

	if request.PageToken != "" {
		// var t time.Time
		//
		// res, err := time.Parse(time.RFC3339Nano, t.String())
		// if err != nil {
		// 	log.Print(err)
		// }
		// tr := filter.EncodeCursor(res, request.GetId())
		//

		println(request.PageToken)
		time, uid, err := filter.DecodeCursor(request.PageToken)
		if err != nil {
			// rtn.LastTicketCursor = errors.New("invalid-cursor").Error()
			print(" ERROR: ", err.Error())
		}
		rtn.LastTicketID = uid
		rtn.LastTicketTimestamp = time
		rtn.LastTicketCursor = request.PageToken
		//
		log.Println("====rtn.LastTicketID====")
		log.Println(rtn.LastTicketID)
		log.Println(rtn.LastTicketTimestamp)
		log.Println(rtn.LastTicketCursor)
		//

	}

	return rtn
}
