package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getPingResponse() gin.H {
	response := gin.H{"success": true}
	return response
}

func PingHandler (c *gin.Context) {
	response := getPingResponse()
	c.JSON(http.StatusOK, response)
}