package config

import (
	"testing"
)


func TestLoadConfig(t *testing.T) {
	goodConfigPath := "../TestResources/test_config.toml"

	expectedGoodConfig := Config{
		Incoming: IncomingData {
			DataSrc: "https://api.hatchways.io/assessment/blog/posts",
		},
		Outgoing: OutgoingData{},
	}

	if goodConfig, err := LoadConfig(goodConfigPath); err != nil {
		t.Error("Failed to load valid config: " + err.Error())
		t.Fail()
	}else {
		if *goodConfig != expectedGoodConfig {
			t.Error("Loaded config did not match expected config")
			t.Fail()
		}
	}
}
