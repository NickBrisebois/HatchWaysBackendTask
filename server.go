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

	fmt.Println(apiConfig)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

}
