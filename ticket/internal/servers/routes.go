package servers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) initHTTPPingRouter() chi.Router {
	return s.chi.Route("/ping", func(r chi.Router) {
		s.logger.Info("sub-router ping: enabled")
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
				s.logger.Errorf("Unable to parse config: %v", err)
			}

			w.Write(cfg)
		})
	})
}
