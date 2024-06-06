package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lesion45/pinterest-clone/internal/http-server/handlers"
	"github.com/lesion45/pinterest-clone/storage/postgres"
	"log/slog"
)

func AddRoutes(router *gin.Engine, log *slog.Logger, s *postgres.Storage) *gin.RouterGroup {
	v1 := router.Group("/v1")

	v1.POST("pin/add", handlers.AddPin(log, s))
	v1.GET("pin/get", handlers.GetPin(log, s))
	v1.GET("pin/get-all", handlers.GetAllPins(log, s))
	v1.POST("pin/delete", handlers.DeletePin(log, s))

	v1.POST("user/register", handlers.Register(log, s))
	v1.POST("user/login", handlers.Login(log, s))

	return v1
}
