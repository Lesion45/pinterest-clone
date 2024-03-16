package storage

import (
	"database/sql"
	"errors"
)

type Storage struct {
	DB *sql.DB
}

var (
	ErrUserExists      error = errors.New("models: user already exists")
	ErrUserNotFound    error = errors.New("models: user not found")
	ErrInvalidPassword error = errors.New("models: invalid password")
	ErrPinNotFound     error = errors.New("models: pin not found")
)
