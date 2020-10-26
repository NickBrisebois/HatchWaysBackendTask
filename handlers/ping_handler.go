package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(c *gin.Context) {
	response := gin.H{"success": true}
	c.JSON(http.StatusOK, response)
}