package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	name := c.Param("name")
	greeting := c.DefaultQuery("greeting", "Hello")

	c.JSON(http.StatusOK, gin.H{
		"message": greeting + ", " + name,
	})
}