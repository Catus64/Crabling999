package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Catus64/repo-observer/api/internal/gitrepo"
)

type CommitHandler struct {
	Repo *gitrepo.Repo
}

func (h *CommitHandler) List(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if err != nil || limit < 1 || limit > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "limit must be 1-100"})
		return
	}

	commits, err := h.Repo.ListCommits(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not read commits"})
		return
	}
	c.JSON(http.StatusOK, commits)
}
