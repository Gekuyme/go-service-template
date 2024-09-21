package main

import (
	"github.com/rs/zerolog/log"
	"project/configs"
	"project/infrastructure/db"
	"project/infrastructure/router"
	"project/internal/repository"
	"project/internal/usecase"
)

func main() {
	cfg := configs.NewConfig()
	database := db.InitPostgresDB(cfg)

	UserRepo := repository.NewUserRepo(database)
	UserUseCase := usecase.NewUserUsecase(UserRepo)

	api := router.InitRouter(UserUseCase)

	log.Info().Msg("Server is running on port 8090")
	api.Run(":8090")

}
