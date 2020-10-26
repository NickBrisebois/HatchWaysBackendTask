package main

import (
	"flag"
	"github.com/NickBrisebois/HatchWaysAppBackend/config"
	"github.com/NickBrisebois/HatchWaysAppBackend/handlers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func serve(config *config.Config) {
	router := gin.Default()

	// Set the config in our handlers to give them access to server configuration
	handlers.SetConfig(config)

	// Initialize our routes to point to our handlers
	api := router.Group(config.Server.APIPrefix)
	api.GET("/ping", handlers.PingHandler)
	api.GET("/posts", handlers.PostsHandler)

	// Configure the HTTP server
	server := &http.Server {
		Addr: config.Server.Address,
		Handler: router,
	}

	// Start the HTTP server
	log.Println("Starting HatchWays API Server")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error starting HatchWays API Server: " + err.Error())
	}

}

func main() {
	configPath := flag.String("config", "./config.toml", "Path to config.toml file")
	flag.Parse()

	apiConfig, err := config.LoadConfig(*configPath)

	if err != nil {
		log.Fatal("Config loading error: " + err.Error())
		return
	}

	serve(apiConfig)
}
