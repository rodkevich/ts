package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"

	"github.com/rodkevich/ts/api/config"
	"github.com/rodkevich/ts/api/pkg/grpc_client"
	"github.com/rodkevich/ts/api/pkg/logger"
)

// Server ...
type Server struct {
	router       *chi.Mux
	logger       logger.Logger
	cfg          *config.Config
	pgConnection *pgxpool.Pool
}

// NewServer ...
func NewServer(logger logger.Logger, cfg *config.Config, pgxPool *pgxpool.Pool) *Server {
	return &Server{
		logger:       logger,
		cfg:          cfg,
		pgConnection: pgxPool,
		router:       chi.NewMux(),
	}
}

var grpcClient *grpc.ClientConn

// Run start a new server
func (s *Server) Run() error {

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	conn, err := grpc_client.NewGRPCClientServiceConn(ctx, "0.0.0.0:5001")
	if err != nil {
		return err
	}

	grpcClient = conn

	// Http //
	s.initHTTPPingRouter()
	s.initHTTPTicketRouter()

	serverHTTP := http.Server{
		Addr:    "0.0.0.0" + s.cfg.HttpServer.Port,
		Handler: s.router,
		// MaxHeaderBytes: maxHeaderBytes,
	}

	serverCtx, serverStopCancel := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig
		// Shutdown signal with grace period
		shutdownCtx, _ := context.WithTimeout(serverCtx, 2*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				s.logger.Fatal("graceful shutdown timed out.. forcing exit")
			}
		}()
		// Trigger graceful shutdown
		err := serverHTTP.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}

		serverStopCancel()

		s.logger.Info("Server Exited Properly")
	}()
	// Run http server
	s.logger.Infof("HTTP Server is listening on port: %v", s.cfg.HttpServer.Port)
	err = serverHTTP.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
	// Wait for context to be stopped
	<-serverCtx.Done()

	return nil
}
