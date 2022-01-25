package config

import (
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	GRPCServer GRPCServer
	HttpServer HttpServer
	Postgres   PostgresConfig
	Logger     Logger
	Redis      RedisConfig
}

type Grpc struct {
	TicketServicePort string
}

type HttpServer struct {
	Port              string
	PprofPort         string
	Timeout           time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	CookieLifeTime    int
	SessionCookieName string
}

// GRPCServer config
type GRPCServer struct {
	AppVersion             string
	Port                   string
	CookieLifeTime         int
	CsrfExpire             int
	SessionID              string
	SessionExpire          int
	Mode                   string
	SessionPrefix          string
	CSRFPrefix             string
	Timeout                time.Duration
	ReadTimeout            time.Duration
	WriteTimeout           time.Duration
	MaxConnectionIdle      time.Duration
	MaxConnectionAge       time.Duration
	SessionGrpcServicePort string
}

// Logger config
type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
	OutputPaths       string
}

// PostgresConfig config
type PostgresConfig struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbname   string
	PostgresqlSSLMode  string
	PgDriver           string
}

// RedisConfig config
type RedisConfig struct {
	RedisAddr      string
	RedisPassword  string
	RedisDB        string
	RedisDefaultDB string
	MinIdleConn    int
	PoolSize       int
	PoolTimeout    int
	Password       string
	DB             int
}

// LoadConfig loads file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func GetConfig(configPath string) (*Config, error) {
	cfgFile, err := LoadConfig(configPath)
	if err != nil {
		return nil, err
	}

	cfg, err := ParseConfig(cfgFile)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "./config/config-docker"
	}
	return "./config/config-local"
}
