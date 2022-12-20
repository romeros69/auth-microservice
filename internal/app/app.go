package app

import (
	"auth-microservice/config"
	v1 "auth-microservice/internal/controller/http/v1"
	"auth-microservice/internal/usecase"
	"auth-microservice/internal/usecase/repo"
	"auth-microservice/pkg/httpserver"
	"auth-microservice/pkg/mongodb"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const collectionName = "accounts"

func Run(cfg *config.Config) {

	mng, err := mongodb.NewMongo(cfg)

	if err != nil {
		log.Fatal("Cannot connect to Mongo")
	}

	handler := gin.New()
	handler.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://51.250.108.95:9000"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Access-Control-Allow-Origin", "Content-Type", "Access-Control-Allow-Credentials", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	userRepo := repo.NewUserRepo(mng, collectionName)
	userUC := usecase.NewUserUseCase(userRepo)
	jwtUC := usecase.NewJwtUseCase(userUC, cfg.JwtSecret)

	v1.NewRouter(handler, userUC, jwtUC)

	serv := httpserver.New(handler, httpserver.Port(cfg.Port))
	interruption := make(chan os.Signal, 1)
	signal.Notify(interruption, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interruption:
		log.Printf("signal: " + s.String())
	case err = <-serv.Notify():
		log.Printf("Notify from http server")
	}

	err = serv.Shutdown()
	if err != nil {
		log.Printf("Http server shutdown")
	}
}
