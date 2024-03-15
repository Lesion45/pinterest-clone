package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lesion45/pinterest-clone/internal/config"
	"github.com/lesion45/pinterest-clone/internal/lib/logger/sl"
	"github.com/lesion45/pinterest-clone/storage/models"
	"golang.org/x/crypto/bcrypt"
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

//
// PIN Table
// pin_id(primary key), image_url, username(reference to User Table)
//

func (s *Storage) CreatePin(imgURL string, username string) error {
	const op = "storage.postgres.createPin"

	stmt := "INSERT INTO pins (imgURL, username) VALUES ($1, $2)"

	_, err := s.db.Exec(stmt, imgURL, username)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) GetPin(id int) (*models.Pin, error) {
	const op = "storage.postgres.GetPin"
	pin := &models.Pin{}

	stmt := "SELECT pin_id, image_url, username FROM pins WHERE pin_id = $1"

	row := s.db.QueryRow(stmt, id)

	err := row.Scan(&pin.ID, &pin.ImageURL, &pin.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrPinNotFound
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return pin, nil
}

func (s *Storage) DeletePin(id int) error {
	const op = "storage.postgres.DeletePin"

	stmt := "DELETE from pins WHERE pin_id = $1"

	_, err := s.db.Exec(stmt, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) GetAllPins() ([]*models.Pin, error) {
	const op = "storage.postgres.GetAllPins"

	stmt := "SELECT pin_id, image_url, username FROM pins"

	rows, err := s.db.Query(stmt)
	defer rows.Close()

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	pins := make([]*models.Pin, 0)

	for rows.Next() {
		pin := &models.Pin{}

		err := rows.Scan(&pin.ID, &pin.ImageURL, &pin.Username)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		pins = append(pins, pin)
	}

	return pins, nil
}

func (s *Storage) GetAllPinsFromUser(username string) ([]*models.Pin, error) {
	const op = "storage.postgres.GetAllPinsFromUser"

	stmt := "SELECT pin_id, image_url, username FROM pins WHERE username = $1"

	rows, err := s.db.Query(stmt, username)
	defer rows.Close()

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	pins := make([]*models.Pin, 0)

	for rows.Next() {
		pin := &models.Pin{}

		err := rows.Scan(&pin.ID, &pin.ImageURL, &pin.Username)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		pins = append(pins, pin)
	}

	return pins, nil
}

//
// User Table
// user_id, username, password(encrypted)
//

// TODO: ADD USER
func (s *Storage) AddUser(username string, password string) error {
	const op = "storage.postgres.AddUser"

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	stmt := "INSERT INTO users (username, passwordHash) VALUES ($1, $2)"

	_, err = s.db.Exec(stmt, username, passwordHash)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) ValidatePassword(username string, password string) error {
	const op = "storage.postgres.ValidatePassword"

	user, err := s.GetUser(username)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return models.ErrInvalidPassword
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) GetUser(username string) (*models.User, error) {
	const op = "storage.postgres.GetUser"

	user := &models.User{}

	stmt := "SELECT username, password FROM users WHERE username = $1"

	row := s.db.QueryRow(stmt, username)

	err := row.Scan(&user.Nickname, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrUserNotFound
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
