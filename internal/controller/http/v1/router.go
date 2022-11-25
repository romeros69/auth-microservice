package v1

import (
	"auth-microservice/internal/usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter(
	handler *gin.Engine,
	userUC usecase.UserContract,
	jwtUC usecase.JwtContract,
) {

	h := handler.Group("/v1")

	{
		newLoginRoutes(h, userUC, jwtUC)
	}
}
