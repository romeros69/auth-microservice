package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Port         string `env:"APP_PORT" envDefault:"9000"`
	JwtSecret    string `env:"SECRET" envDefault:"LEXA_KUTSENKA"`
	LogLevel     string `env:"LOG_LEVEL" envDefault:"error"`
	MongoURL     string `env:"MONGO_RUL" envDefault:"mongodb://mongo:mongo@mongodb:27017"`
	MongoDB      string `env:"MONGO_DB" envDefault:"users"`
	GrpcProtocol string `env:"GRPC_PROT" envDefault:"tcp"`
	GrpcURL      string `env:"GRPC_URL" envDefault:":9000"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
