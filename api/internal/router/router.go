package router

import (
	"github.com/gin-gonic/gin"

	"github.com/Catus64/repo-observer/api/internal/gitrepo"
	"github.com/Catus64/repo-observer/api/internal/handlers"
)

func New(repo *gitrepo.Repo) *gin.Engine {
	r := gin.Default()

	r.GET("/health", handlers.Health)
	r.GET("/hello/:name", handlers.Hello)

	commits := &handlers.CommitHandler{Repo: repo}
	api := r.Group("/api/v1")
	api.GET("/commits", commits.List)

	return r
}
