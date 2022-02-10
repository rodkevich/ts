package main

import (
	sysLog "log"
	"os"

	"github.com/rodkevich/ts/api/config"
	"github.com/rodkevich/ts/api/internal/server"
	"github.com/rodkevich/ts/api/pkg/logger"
	"github.com/rodkevich/ts/api/pkg/redis"
	// "github.com/rodkevich/ts/api/pkg/postgres"
)

func main() {
	configPath := config.GetConfigPath(os.Getenv("CONFIG"))
	cfg, err := config.GetConfig(configPath)

	if err != nil {
		sysLog.Fatalf("Loading config: %v", err)
	}

	// logger
	log := logger.New(cfg)
	log.InitLogger()

	log.Info("Starting tickets server")
	log.Infof(
		"AppVersion: %s, LogLevel: %s, Mode: %s",
		cfg.GRPCServer.AppVersion,
		cfg.Logger.Level,
		cfg.GRPCServer.Mode,
	)
	log.Infof("Success parsed config: %#v", cfg.GRPCServer.AppVersion)

	// // database
	// pgxConn, err := postgres.NewPgxConn(cfg)
	// if err != nil {
	// 	log.Fatal("cannot connect to postgres", err)
	// }
	// defer pgxConn.Close()
	//
	// log.Infof("%-v", pgxConn.Config().ConnString())

	// using redis for cache
	redisClient := redis.NewRedisClient(cfg)
	log.Infof("Redis connected: %-v", redisClient.PoolStats())

	// server
	s := server.NewServer(log, cfg, redisClient, nil)
	log.Fatal(" server stopped running: ERRORS: ", s.Run())
}
