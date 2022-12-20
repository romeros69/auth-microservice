package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"strings"
)

type Config struct {
	Port         string `env:"APP_PORT" envDefault:"9000"`
	JwtSecret    string `env:"SECRET" envDefault:"LEXA_KUTSENKA"`
	LogLevel     string `env:"LOG_LEVEL" envDefault:"error"`
	MongoURL     string
	MongoDB      string `env:"MONGO_DB" envDefault:"users"`
	GrpcProtocol string `env:"GRPC_PROT" envDefault:"tcp"`
	GrpcURL      string `env:"GRPC_URL" envDefault:":9011"`
}

func NewConfig() (*Config, error) {
	const DB_NAME = "users"
	DB_HOSTS := []string{
		"rc1b-t408fzercgkkda2o.mdb.yandexcloud.net:27018",
	}
	const DB_USER = "romich"
	const DB_PASS = "romichevgen"

	const CACERT = "/home/romich-v2/.mongodb/root.crt"
	url := fmt.Sprintf("mongodb://%s:%s@%s/%s?tls=true&tlsCaFile=%s",
		DB_USER,
		DB_PASS,
		strings.Join(DB_HOSTS, ","),
		DB_NAME,
		CACERT)
	cfg := &Config{}
	cfg.MongoURL = url
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
