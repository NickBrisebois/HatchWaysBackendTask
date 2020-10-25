package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Incoming IncomingData `toml:"incoming_data"`
	Outgoing OutgoingData `toml:"outgoing_data"`
}

type IncomingData struct {
	DataSrc string `toml:"data_src"`
}

type OutgoingData struct {

}

func LoadConfig (configPath string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		return nil, err
	}

	return &config, nil
}