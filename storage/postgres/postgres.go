package postgres

import (
	"database/sql"
	"fmt"
	"github.com/lesion45/pinterest-clone/internal/config"
	"github.com/lesion45/pinterest-clone/internal/lib/logger/sl"
	"log/slog"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(cfg config.Config, log slog.Logger) (*Storage, error) {
	const op = "storage.postgres.NewStorage"
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", cfg.DB.Username, cfg.DB.DBName, cfg.DB.Password, cfg.DB.Host, cfg.DB.SSLMode)

	db, err := sql.Open("postgres", connStr)

	stmt := "CREATE TABLE IF NOT EXISTS users (user_id serial PRIMARY KEY, username varchar(255) UNIQUE NOT NULL, password char(60) NOT NULL)"

	_, err = db.Exec(stmt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt = "CREATE TABLE IF NOT EXISTS pins (pin_id serial PRIMARY KEY, image_url varchar(255) NOT NULL, username varchar(255) NOT NULL REFERENCES users (username))"

	_, err = db.Exec(stmt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err != nil {
		log.Error("failed to initialize storage", sl.Err(err))
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Error("storage doesn't response", sl.Err(err))
		return nil, err
	}

	return &Storage{db: db}, nil
}

func (s *Storage) createPin(imgURL, username string) error {
	stmt := "INSERT INTO pins (imgURL, username) VALUES ($1, $2)"

	_, err := s.db.Exec(stmt, imgURL, username)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) getPin(id int) (int, error) {
	return 1, nil
}
