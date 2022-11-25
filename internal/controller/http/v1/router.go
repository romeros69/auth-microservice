package v1

import (
	_ "auth-microservice/docs"
	"auth-microservice/internal/usecase"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(
	handler *gin.Engine,
	userUC usecase.UserContract,
	jwtUC usecase.JwtContract,
) {

	handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	h := handler.Group("/v1")

	{
		newLoginRoutes(h, userUC, jwtUC)
		newLogoutRoutes(h, jwtUC)
	}
}
