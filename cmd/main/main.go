package main

import (
	"auth-microservice/config"
	"auth-microservice/internal/app"
	"fmt"
	"log"
)

// @tittle auth
// @version 1.0
// @description API for auth microservice

// @host localhost:9000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("Error in parsing config")
	}
	fmt.Println("hello")
	app.Run(cfg)
}
