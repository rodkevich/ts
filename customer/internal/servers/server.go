package servers

import (
	"context"
	"encoding/json"
	pb "github.com/rodkevich/ts/customer/proto/customer/v1"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/go-playground/validator/v10"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	"github.com/go-chi/chi/v5"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/rodkevich/ts/customer/config"
	handlers "github.com/rodkevich/ts/customer/internal/blueprints/grpc"
	"github.com/rodkevich/ts/customer/internal/controllers"
	"github.com/rodkevich/ts/customer/internal/repositories/postgres"
	"github.com/rodkevich/ts/customer/pkg/logger"
)

const (
	certFile              = "ssl/server.crt"
	keyFile               = "ssl/server.pem"
	maxHeaderBytes        = 1 << 20
	customerCachePrefix   = "customers:"
	customerCacheDuration = time.Minute * 15
)

// Server ...
type Server struct {
	chi          *chi.Mux
	logger       logger.Logger
	cfg          *config.Config
	pgConnection *pgxpool.Pool
}

// NewServer ...
func NewServer(
	logger logger.Logger,
	cfg *config.Config,
	pgxPool *pgxpool.Pool,
) *Server {
	return &Server{logger: logger, cfg: cfg, pgConnection: pgxPool, chi: chi.NewMux()}
}

func (s *Server) Run() error {

	// Validator //
	_ = validator.New()

	// DEBUG //
	go func() {
		s.logger.Infof("Starting Debug Server on PORT: %s", s.cfg.HttpServer.PprofPort)
		if err := http.ListenAndServe(s.cfg.HttpServer.PprofPort, http.DefaultServeMux); err != nil {
			s.logger.Errorf("Error PPROF ListenAndServe: %s", err)
		}
	}()

	// Grpc //
	lis, err := net.Listen("tcp", s.cfg.GRPCServer.Port)
	if err != nil {
		return err
	}
	defer lis.Close()

	serverGRPC := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: s.cfg.GRPCServer.MaxConnectionIdle * time.Minute,
		Timeout:           s.cfg.GRPCServer.Timeout * time.Second,
		MaxConnectionAge:  s.cfg.GRPCServer.MaxConnectionAge * time.Minute,
		Time:              s.cfg.GRPCServer.Timeout * time.Minute,
	}),
		grpc.ChainUnaryInterceptor(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpcrecovery.UnaryServerInterceptor(),
			// pls.logger,
		),
	)

	customerDB := postgres.NewCustomer(s.pgConnection)
	customerController := controllers.NewCustomerController(s.logger, customerDB)
	customerService := handlers.NewCustomerGrpcService(s.logger, customerController)
	pb.RegisterCustomerServiceServer(serverGRPC, customerService)

	go func() {
		s.logger.Infof("GRPC Server is listening on port: %v", s.cfg.GRPCServer.Port)
		err := serverGRPC.Serve(lis)
		if err != nil {
			s.logger.Fatal("serverGRPC.Serve(lis) is not running")
		}
	}()

	if s.cfg.GRPCServer.Mode != "Production" {
		reflection.Register(serverGRPC)
	}

	// Http //
	s.chi.Route("/ping", func(r chi.Router) {
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
	serverHTTP := http.Server{
		Addr:           "0.0.0.0:3333",
		Handler:        s.chi,
		MaxHeaderBytes: maxHeaderBytes,
	}

	// Listen for syscall signals for process to interrupt/quit
	serverCtx, serverStopCancel := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig
		// Shutdown signal with grace period of 30 seconds
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

		serverGRPC.GracefulStop()
		serverStopCancel()

		s.logger.Info("Server Exited Properly")
	}()
	// Run http server
	err = serverHTTP.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
	// Wait for context to be stopped
	<-serverCtx.Done()

	return nil
}
