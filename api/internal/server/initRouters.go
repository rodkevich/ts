package server

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"google.golang.org/grpc"

	"github.com/rodkevich/ts/api/internal/util"
	"github.com/rodkevich/ts/api/pkg/filter"
	v1 "github.com/rodkevich/ts/api/proto/ticket/v1"
)

func (s *Server) initHTTPPingRouter() chi.Router {

	return s.router.Route("/ping", func(r chi.Router) {
		s.logger.Info("sub-router 'api-service: /ping': enabled")

		r.Use(middleware.RequestID)
		r.Use(middleware.Logger)
		r.Use(middleware.RealIP)
		r.Use(middleware.Recoverer)
		r.Use(render.SetContentType(render.ContentTypeJSON))

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			render.JSON(w, r, "pong")

		})

		r.Get("/config", func(w http.ResponseWriter, r *http.Request) {
			// cfg, err := json.Marshal(s.cfg)
			// if err != nil {
			// 	s.logger.Errorf("api-service: /config : Unable to parse config: %v", err)
			// }

			render.JSON(w, r, s.cfg)
		})
	})
}

func (s *Server) initHTTPTicketRouter() chi.Router {

	return s.router.Route("/ticket", func(r chi.Router) {
		s.logger.Info("sub-router 'api-service: /ticket': enabled")

		r.Use(middleware.RequestID)
		r.Use(middleware.Logger)
		r.Use(middleware.RealIP)
		r.Use(middleware.Recoverer)
		r.Use(render.SetContentType(render.ContentTypeJSON))

		r.Get("/{ticketID}", func(w http.ResponseWriter, r *http.Request) {
			var (
				opts   []grpc.CallOption
				ti     = v1.NewTicketServiceClient(grpcClient)
				ctx    = context.Background()
				ticket *v1.ListTicketsResponse
			)

			if ticketID := chi.URLParam(r, "ticketID"); ticketID != "" {
				tr := &v1.GetTicketRequest{
					Id: ticketID,
				}

				getResp, err := ti.GetTicket(ctx, tr, opts...)
				ticket = getResp
				if err != nil {
					s.logger.Error(err.Error())
				}
			}
			render.JSON(w, r, ticket)

			// resp, _ := json.Marshal(&ticket)
			// w.Write(resp)
		})

		r.Get("/list", func(w http.ResponseWriter, r *http.Request) {
			baseFilter := filter.NewFromURL(r.URL.Query())
			id := r.URL.Query().Get("id")
			ticketFields, _ := util.FieldsFromURL(r.URL.Query(), "ticket")

			req := v1.ListTicketsRequest{
				Id:        id,
				Reverse:   baseFilter.Reversed,
				Search:    baseFilter.Search,
				Paging:    baseFilter.Paging,
				PageSize:  uint32(baseFilter.Size),
				PageToken: "",
				Fields:    ticketFields,
				// Fields: FieldsFromURL(r),
			}

			resp, err := json.Marshal(&req)
			if err != nil {
				s.logger.Errorf("api-service: /ticket/list: %v", err)
			}

			w.Write(resp)

		})

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			err := r.ParseForm()
			if err != nil {
				s.logger.Errorf("api-service: /ticket : Unable to parse request: %v", err)
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					s.logger.Errorf("api-service: io.ReadCloser: /ticket : Unable to parse request: %v", err)
				}
			}(r.Body)
			buf := bufio.NewScanner(r.Body)
			for buf.Scan() {
				fmt.Println(buf.Text())
			}
			w.Write([]byte("form parsed"))
		})
	})
}
