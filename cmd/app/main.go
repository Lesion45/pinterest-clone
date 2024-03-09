package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lesion45/pinterest-clone/internal/config"
	rts "github.com/lesion45/pinterest-clone/internal/http-server/routes"
	"github.com/lesion45/pinterest-clone/internal/lib/logger/sl"
	"github.com/lesion45/pinterest-clone/storage/postgres"
	_ "github.com/lib/pq"
	"log/slog"
	"net/http"
	"os"
)

type App struct {
	server  *http.Server
	router  *gin.Engine
	storage *postgres.Storage
	log     *slog.Logger
}

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log = log.With(slog.String("env", cfg.Env))

	log.Info("initializing server", slog.String("address", cfg.Server.Address))
	log.Debug("logger debug mode enabled")

	_, err := postgres.NewStorage(*cfg, *log)
	if err != nil {
		log.Error("database initialization error", sl.Err(err))
		os.Exit(1)
	}
	log.Info("Storage is active")

	router := gin.Default()

	routes := rts.AddRoutes(router)
	if routes == nil {
		log.Error("Routes error", sl.Err(err))
		os.Exit(1)
	}

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
