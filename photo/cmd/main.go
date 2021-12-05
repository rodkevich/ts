package main

import (
	"log"
	"os"

	"github.com/rodkevich/ts/photo/config"
	"github.com/rodkevich/ts/photo/internal/servers"
	"github.com/rodkevich/ts/photo/pkg/logger"
	"github.com/rodkevich/ts/photo/pkg/postgres"
)

func main() {
	configPath := config.GetConfigPath(os.Getenv("config"))
	cfg, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatalf("Loading config: %v", err)
	}

	// logger
	appLogger := logger.New(cfg)
	appLogger.InitLogger()
	appLogger.Info("Starting user server")
	appLogger.Infof(
		"AppVersion: %s, LogLevel: %s, Mode: %s",
		cfg.GRPCServer.AppVersion,
		cfg.Logger.Level,
		cfg.GRPCServer.Mode,
	)
	appLogger.Infof("Success parsed config: %#v", cfg.GRPCServer.AppVersion)
	// database
	pgxConn, err := postgres.NewPgxConn(cfg)
	if err != nil {
		appLogger.Fatal("cannot connect to postgres", err)
	}
	defer pgxConn.Close()

	appLogger.Infof("%-v", pgxConn.Config().ConnString())
	// server
	s := servers.NewServer(appLogger, cfg, pgxConn)
	appLogger.Fatal(" server stopped running: ERRORS: ", s.Run())
}
