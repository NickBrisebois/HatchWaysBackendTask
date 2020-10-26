package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Generates a ping response in a testable way
func getPingResponse() gin.H {
	response := gin.H{"success": true}
	return response
}

// PingHandler handles answering ping requests
func PingHandler (c *gin.Context) {
	response := getPingResponse()
	c.JSON(http.StatusOK, response)
}