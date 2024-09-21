package usecase

import (
	"github.com/rs/zerolog/log"
	"project/internal/domain"
	"project/internal/repository"
)

type UserUsecase struct {
	userRepo *repository.UserRepo
}

func NewUserUsecase(userRepo *repository.UserRepo) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (u *UserUsecase) CreateUser(user *domain.User) error {
	err := u.userRepo.CreateUser(user)
	if err != nil {
		log.Error().Err(err).Msg("Error creating user")
		return err
	}
	return nil
}
