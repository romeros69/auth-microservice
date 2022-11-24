package main

import (
	"auth-microservice/config"
	"auth-microservice/internal/app"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("Error in parsing config")
	}
	fmt.Println("hello")
	app.Run(cfg)
}
