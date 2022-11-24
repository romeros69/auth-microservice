package main

import (
	"auth-microservice/config"
	"fmt"
	"log"
)

func main() {
	_, err := config.NewConfig()
	if err != nil {
		log.Fatal("Error in parsing config")
	}
	fmt.Println("hello")
	//app.Run(cfg)
}
