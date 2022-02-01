package servers

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4/pgxpool"

	ticketCFG "github.com/rodkevich/ts/ticket/config"
	"github.com/rodkevich/ts/ticket/internal/controllers/ticket"
	ticketGRPCHandlers "github.com/rodkevich/ts/ticket/internal/handlers/ticket/grpc"
	ticketRepoPG "github.com/rodkevich/ts/ticket/internal/repositories/ticket/postgres"
	ticketLogger "github.com/rodkevich/ts/ticket/pkg/logger"
	ticketProto "github.com/rodkevich/ts/ticket/proto/ticket/v1"
)

const (
	certFile            = "ssl/server.crt"
	keyFile             = "ssl/server.pem"
	maxHeaderBytes      = 1 << 20
	ticketCachePrefix   = "tickets:"
	ticketCacheDuration = time.Minute * 15
)

// Server ...
type Server struct {
	chi          *chi.Mux
	logger       ticketLogger.Logger
	cfg          *ticketCFG.Config
	pgConnection *pgxpool.Pool
}

// NewServer ...
func NewServer(logger ticketLogger.Logger, cfg *ticketCFG.Config, pgxPool *pgxpool.Pool) *Server {
	return &Server{
		logger:       logger,
		cfg:          cfg,
		pgConnection: pgxPool,
		chi:          chi.NewMux(),
	}
}

func (s *Server) Run() error {

	// Validator //
	validate := validator.New()

	// DEBUG //
	go func() {
		s.logger.Infof("Starting Debug Server on PORT: %s", s.cfg.HttpServer.PprofPort)
		if err := http.ListenAndServe(s.cfg.HttpServer.PprofPort, http.DefaultServeMux); err != nil {
			s.logger.Errorf("Error PPROF ListenAndServe: %s", err)
		}
	}()

	// Grpc //
	listener, err := net.Listen("tcp", s.cfg.GRPCServer.Port)
	if err != nil {
		return err
	}
	defer listener.Close()

	srv := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: s.cfg.GRPCServer.MaxConnectionIdle * time.Minute,
		Timeout:           s.cfg.GRPCServer.Timeout * time.Second,
		MaxConnectionAge:  s.cfg.GRPCServer.MaxConnectionAge * time.Minute,
		Time:              s.cfg.GRPCServer.Timeout * time.Minute,
	}),
		grpc.ChainUnaryInterceptor(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
			// pls.ticketLogger,
		),
	)

	// Tickets-service //
	r := ticketRepoPG.New(s.pgConnection)
	c := ticket.New(s.logger, r)
	h := ticketGRPCHandlers.New(s.logger, c, validate)
	ticketProto.RegisterTicketServiceServer(srv, h)

	go func() {
		s.logger.Infof("GRPC Server is listening on port: %v", s.cfg.GRPCServer.Port)
		err := srv.Serve(listener)
		if err != nil {
			s.logger.Fatal("srv.Serve(listener) is not running")
		}
	}()

	if s.cfg.GRPCServer.Mode != "Production" {
		reflection.Register(srv)
	}

	// Http //
	s.initHTTPPingRouter()

	serverHTTP := http.Server{
		Addr:           "0.0.0.0" + s.cfg.HttpServer.Port,
		Handler:        s.chi,
		MaxHeaderBytes: maxHeaderBytes,
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

		srv.GracefulStop()
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
