package models

import "errors"

var (
	ErrUserExists      error = errors.New("models: user already exists")
	ErrUserNotFound    error = errors.New("models: user not found")
	ErrInvalidPassword error = errors.New("models: invalid password")
	ErrPinNotFound     error = errors.New("models: pin not found")
)

type User struct {
	ID       int
	Nickname string
	Login    string
	Password []byte
}

type Pin struct {
	ID       int
	ImageURL string
	Username string
}
