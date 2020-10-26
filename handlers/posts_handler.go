package handlers

import (
	"errors"
	"fmt"
	"github.com/NickBrisebois/HatchWaysAppBackend/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var serverConfig *config.Config

type queryConfig struct {
	tags []string
	sortBy string
	direction string
}

// SetConfig is a Get function for setting handler package's serverConfig
func SetConfig(configToSet *config.Config) {
	serverConfig = configToSet
}

// PostsHandler handles GET requests for retrieving and sorting posts
func PostsHandler (c *gin.Context) {
	queryConfig, err := getAndValidateQueries(c)
	if err != nil {
		errMsg := gin.H{"Bad Request": err.Error()}
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	posts, err := getCombinedPosts(queryConfig.tags)
	if err != nil {
		errMsg := gin.H{"Error getting posts": err.Error()}
		c.JSON(http.StatusInternalServerError, errMsg)
		return
	}

	fmt.Println(posts)
}

// isValidSortBy checks the selected sortBy method and returns true or false whether its valid or not
func isValidSortBy(chosenSort string) bool {
	for _, sort := range serverConfig.APISettings.AcceptableSortBy {
		if chosenSort == sort {
			return true
		}
	}
	return false
}

// getAndValidateQueries checks the given GET queries and error checks, then defaults if needed then returns values
func getAndValidateQueries(c *gin.Context) (*queryConfig, error) {
	requestFields := c.Request.URL.Query()
	tags := requestFields.Get("tags")
	sortBy := requestFields.Get("sortBy")
	direction := requestFields.Get("direction")

	if tags == "" {
		// If tags is not provided then we error out. It's a mandatory field
		return nil, errors.New("tags parameter is required")
	}

	if sortBy == "" {
		// If sort by is not provided, we default to ID
		sortBy = "id"
	}else if !isValidSortBy(sortBy){
		// If sort by method is not valid, return an error
		return nil, errors.New("sortBy parameter is invalid")
	}

	if direction == "" {
		// If direction is not provided, we default to ascending
		direction = "asc"
	}else if direction != "asc" && direction != "desc" {
		// if direction is not either of the two valid ones, return an error
		return nil, errors.New("direction parameter is invalid")
	}

	return &queryConfig{
		tags: strings.Split(tags, ","),
		sortBy: sortBy,
		direction: direction,
	}, nil
}