package v1

import (
	"auth-microservice/internal/entity"
	"auth-microservice/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginRoutes struct {
	userUC usecase.UserContract
	jwtUC  usecase.JwtContract
}

func newLoginRoutes(handler *gin.RouterGroup, u usecase.UserContract, j usecase.JwtContract) {
	lg := loginRoutes{userUC: u, jwtUC: j}

	handler.POST("/login", lg.login)
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *loginRoutes) login(c *gin.Context) {
	var lg loginRequest
	if err := c.ShouldBindJSON(&lg); err != nil {
		errorResponse(c, http.StatusBadRequest, "cannot parse user data")
		return
	}
	err := validateLogin(lg.Email, lg.Password)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "error format of login")
		return
	}
	err = l.jwtUC.CompareUserPassword(c.Request.Context(), entity.User{
		Email:    lg.Email,
		Password: lg.Password,
	})
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "wrong password")
		return
	}
	token, err := l.jwtUC.GenerateToken(lg.Email)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "cannot generate token")
		return
	}
	c.Header("Authorization", "Bearer: "+token)
	c.JSON(http.StatusOK, nil)
}
