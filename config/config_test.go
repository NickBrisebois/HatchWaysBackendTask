package config

import (
	"testing"
)

const testConfig = "../TestResources/test_config.toml"

func TestLoadConfig(t *testing.T) {
	expectedGoodConfig := Config{
		Incoming: IncomingData {
			DataSrc: "https://api.hatchways.io/assessment/blog/posts",
		},
		Server: Server{
			APIPrefix: "/api",
			Address: "localhost:8080",
		},
	}

	if goodConfig, err := LoadConfig(testConfig); err != nil {
		t.Error("Failed to load valid config: " + err.Error())
		t.Fail()
	}else {
		if *goodConfig != expectedGoodConfig {
			t.Error("Loaded config did not match expected config")
			t.Fail()
		}
	}
}
