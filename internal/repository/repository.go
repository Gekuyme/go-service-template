package repository

import (
	"database/sql"
	"github.com/rs/zerolog/log"
	"project/internal/domain"
	"strconv"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) CreateUser(user *domain.User) error {
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`

	res, err := r.DB.Exec(query, user.Name, user.Email, user.Password)

	if err != nil {
		log.Error().Err(err).Msg("Error creating user")
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Error().Err(err).Msg("Error getting last insert id")
		return err
	}
	log.Info().Msg(strconv.FormatInt(id, 10))
	return nil
}
