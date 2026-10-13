package router

import (
	"github.com/gin-gonic/gin"

	"github.com/Catus64/repo-observer/api/internal/handlers"
)

func New() *gin.Engine {
	r := gin.Default()

	r.GET("/health", handlers.Health)
	r.GET("/hello/:name", handlers.Hello)
	r.GET("/ws", handlers.Echo)
	return r
}
