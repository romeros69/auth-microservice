package main

import (
	"auth-microservice/config"
	"auth-microservice/internal/app"
	"log"
)

// @tittle auth
// @version 1.0
// @description API for auth microservice

// @host 51.250.108.95:9000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("Error in parsing config")
	}

	app.Run(cfg)
}
