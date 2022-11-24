package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Port         string `env:"APP_PORT" envDefault:"9000"`
	JwtSecret    string `env:"SECRET" envDefault:"LEXA_KUTSENKA"`
	LogLevel     string `env:"LOG_LEVEL" envDefault:"error"`
	PostgresUrl  string `env:"POSTGRES_URL" envDefault:"postgresql://postgres1488:postgres1488@psql_auth:5432/hachi"`
	GrpcProtocol string `env:"GRPC_PROT" envDefault:"tcp"`
	GrpcURL      string `env:"GRPC_URL" envDefault:":9011"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
