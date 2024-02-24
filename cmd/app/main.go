package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lesion45/pinterest-clone/internal/config"
	"github.com/lesion45/pinterest-clone/internal/lib/logger/sl"
	_ "github.com/lib/pq"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log = log.With(slog.String("env", cfg.Env))

	log.Info("initializing server", slog.String("address", cfg.Server.Address))
	log.Debug("logger debug mode enabled")

	// TODO: replace with a single function
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", cfg.DB.Host, cfg.DB.Username, cfg.DB.DBName, cfg.DB.Password, cfg.DB.Host, cfg.DB.SSLMode)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Error("failed to initialize storage", sl.Err(err))
		os.Exit(1)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Error("storage doesn't response", sl.Err(err))
		os.Exit(1)
	}

	router := mux.NewRouter()

	server := &http.Server{
		Addr:         cfg.Server.Address,
		Handler:      router,
		ReadTimeout:  cfg.Server.TimeOut,
		WriteTimeout: cfg.Server.TimeOut,
		IdleTimeout:  cfg.Server.IdleTimeOut,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}
	log.Error("server stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "dev":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "prod":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
