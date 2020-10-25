package main

import (
	"flag"
	"fmt"
	"github.com/NickBrisebois/HatchWaysAppBackend/config"
	"log"
)

func main() {
	configPath := flag.String("config", "./config.toml", "Path to config.toml file")
	flag.Parse()

	apiConfig, err := config.LoadConfig(*configPath)

	if err != nil {
		log.Fatal("Config loading error: " + err.Error())
		return
	}

	postsRetriever := NewPostsRetriever(apiConfig)
	posts, err := postsRetriever.GetPosts("tech")
	fmt.Print(posts)

	if err != nil {
		log.Fatal("Error retrieving posts: " + err.Error())
	}
}
