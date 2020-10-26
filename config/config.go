package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Incoming IncomingData `toml:"incoming_data"`
	APISettings APISettings `toml:"api_settings"`
	Server Server `toml:"server"`
}

type IncomingData struct {
	DataSrc string `toml:"data_src"`
}

type APISettings struct {
	AcceptableSortBy []string `toml:"acceptable_sortby"`
}

type Server struct {
	APIPrefix string `toml:"api_prefix"`
	Address string `toml:"address"`
}

func LoadConfig (configPath string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		return nil, err
	}

	return &config, nil
}