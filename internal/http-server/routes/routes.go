package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lesion45/pinterest-clone/internal/http-server/handlers"
	"github.com/lesion45/pinterest-clone/storage/postgres"
	"log/slog"
)

func AddRoutes(router *gin.Engine, log *slog.Logger, str *postgres.Storage) *gin.RouterGroup {
	v1 := router.Group("/v1")

	v1.POST("/create", handlers.Register(log, str))

	return v1
}
