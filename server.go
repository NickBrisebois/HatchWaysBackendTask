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

	api := router.Group(config.Server.APIPrefix)
	api.GET("/ping", handlers.Ping)

	server := &http.Server {
		Addr: config.Server.Address,
		Handler: router,
	}

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
