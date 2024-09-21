package router

import (
	"github.com/gin-gonic/gin"
	"project/internal/delivery/http"
	"project/internal/usecase"
)

func InitRouter(usecase *usecase.UserUsecase) *gin.Engine {
	router := gin.Default()
	http.NewUserHandler(router, usecase)
	return router
}
