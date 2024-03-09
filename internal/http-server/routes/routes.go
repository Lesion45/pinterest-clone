package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lesion45/pinterest-clone/internal/http-server/handlers"
)

func AddRoutes(r *gin.Engine) *gin.RouterGroup {
	v1 := r.Group("/v1")

	v1.POST("/create", handlers.CreatePin)

	return v1
}
