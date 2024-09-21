package db

import (
	"database/sql"
	"fmt"
	"github.com/rs/zerolog/log"
	"project/configs"
)

func InitPostgresDB(cfg *configs.Config) *sql.DB {
	str := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBname)

	db, err := sql.Open("postgres", str)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
		return nil
	}
	return db
}
