package main

import (
	config "github.com/NickBrisebois/HatchWaysAppBackend/config"
	"testing"
	"reflect"
)

const testConfig = "./TestResources/test_config.toml"

func TestPostsRetriever_GetPosts(t *testing.T) {
	config, err := config.LoadConfig(testConfig)

	expectedTestData := Post{
		"Rylee Paul",
		9,
		1,
		960,
		0.13,
		50361,
		[]string{"tech", "health"},
	}
	_ = expectedTestData

	if err != nil {
		t.Error("Error loading test config: " + err.Error())
		t.Fail()
	}

	postRetriever := NewPostsRetriever(config)

	if posts, err := postRetriever.GetPosts("tech"); err != nil {
		t.Error("Error loading posts: " + err.Error())
		t.Fail()
	}else {
		firstPost := posts.Posts[0]
		if !reflect.DeepEqual(firstPost, expectedTestData) {
			t.Error("Retrieved post data is not equal to expected test data")
			t.Fail()
		}
	}
}
