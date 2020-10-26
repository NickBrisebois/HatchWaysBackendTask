package handlers

import (
	"github.com/NickBrisebois/HatchWaysAppBackend/config"
	"reflect"
	"testing"
)

const testConfig = "../TestResources/test_config.toml"

func TestPostsRetriever_GetPosts(t *testing.T) {
	conf, err := config.LoadConfig(testConfig)
	SetConfig(conf)

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
		t.Error("Error loading test conf: " + err.Error())
		t.Fail()
	}

	if posts, err := getPosts("tech"); err != nil {
		t.Error("Error loading posts: " + err.Error())
		t.Fail()
	} else {
		firstPost := posts.Posts[0]
		if !reflect.DeepEqual(firstPost, expectedTestData) {
			t.Error("Retrieved post data is not equal to expected test data")
			t.Fail()
		}
	}
}
