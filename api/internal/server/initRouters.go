package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/rodkevich/ts/api/pkg/filter"
)

func (s *Server) initHTTPPingRouter() chi.Router {
	return s.router.Route("/ping", func(r chi.Router) {
		s.logger.Info("sub-router 'api-service: /ping': enabled")
		r.Use(middleware.RequestID)
		r.Use(middleware.Logger)
		r.Use(middleware.RealIP)
		r.Use(middleware.Recoverer)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})
		r.Get("/config", func(w http.ResponseWriter, r *http.Request) {
			cfg, err := json.Marshal(s.cfg)
			if err != nil {
				s.logger.Errorf("api-service: /config : Unable to parse config: %v", err)
			}

			w.Write(cfg)
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

		r.Get("/:id", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("ticket router"))
		})

		r.Get("/list", func(w http.ResponseWriter, r *http.Request) {
			f := filter.NewFromURL(r.URL.Query())
			resp, err := json.Marshal(f)
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

			w.Write([]byte("form parsed"))
		})
	})
}
