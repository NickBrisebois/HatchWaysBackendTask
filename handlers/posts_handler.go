package handlers

import (
	"github.com/NickBrisebois/HatchWaysAppBackend/config"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var config *config.Config

type queryConfig struct {
	tags []string
	sortBy string
	direction string
}

func SetConfig(configToSet *Config.Config) {
	config = configToSet
}

func PostsHandler (c *gin.Context) {
	queryConfig, err := getAndValidateQueries(c)
	if err != nil {
		errMsg := gin.H{"Bad Request": err.Error()}
		c.JSON(http.StatusBadRequest, errMsg)
	}

	fmt.Println(queryConfig.tags)
	fmt.Println(queryConfig.sortBy)
	fmt.Println(queryConfig.direction)
}

func getAndValidateQueries(c *gin.Context) (*queryConfig, error) {
	requestFields := c.Request.URL.Query()
	tags := requestFields.Get("tags")
	sortBy := requestFields.Get("sortBy")
	direction := requestFields.Get("direction")

	// If tags is not provided then we error out. It's a mandatory field
	if tags == "" {
		return nil, errors.New("tags parameter is required")
	}

	// If sort by is not provided, we default to ID
	if sortBy == "" {
		sortBy = "id"
	}else {

	}

	// If direction is not provided, we default to ascending
	if direction == "" {
		direction = "asc"
	}else if direction != "asc" && direction != "desc" {
		return nil, errors.New("direction parameter is invalid")
	}

	return &queryConfig{
		tags: strings.Split(tags, ","),
		sortBy: sortBy,
		direction: direction,
	}, nil
}