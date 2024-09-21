package http

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"project/internal/domain"
	"project/internal/usecase"
)

type UserHandler struct {
	UserUseCase *usecase.UserUsecase
}

func NewUserHandler(router *gin.Engine, userUseCase *usecase.UserUsecase) {
	handler := UserHandler{UserUseCase: userUseCase}
	router.POST("/user", handler.CreateUser)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var User *domain.User

	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := h.UserUseCase.CreateUser(User)
	if err != nil {
		log.Error().Err(err).Msg("Error creating user")
	}
	c.JSON(200, gin.H{"message": "User created successfully"})
}
