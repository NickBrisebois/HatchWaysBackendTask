// Package handlers provides handlers for different API routes
package handlers

import (
	"github.com/NickBrisebois/HatchWaysAppBackend/config"
	"reflect"
	"testing"
)

const testConfig = "../TestResources/test_config.toml"

// TestGetPosts tests retrieving posts from HatchWays API
func TestGetPosts(t *testing.T) {
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

// getExamplePosts returns an array of posts to be used in testing the sorting of posts. By default
// this function returns posts in ascending order
func getExamplePosts() []Post {
	return []Post{
		Post{
			"First Post",
			1,
			1,
			1,
			0.1,
			1,
			[]string{"tech", "health"},
		},
		Post{
			"Second Post",
			2,
			2,
			2,
			2,
			2,
			[]string{"tech", "health"},
		},
		Post{
			"Third Post",
			3,
			3,
			3,
			3,
			3,
			[]string{"tech", "health"},
		},
	}
}

// TestSortPostsAsc tests sorting posts in ascending manner.
// Sorting should not change order as posts are already in ascending manner
func TestSortPostsAsc(t *testing.T) {
	conf, err := config.LoadConfig(testConfig)
	SetConfig(conf)
	if err != nil {
		t.Error("Error loading test conf: " + err.Error())
		t.Fail()
	}

	sortTypes := conf.APISettings.AcceptableSortBy
	for _, sortType := range sortTypes {
		examplePosts := getExamplePosts()
		sorted := sortPosts(examplePosts, "asc", sortType)
		if !reflect.DeepEqual(sorted, examplePosts) {
			t.Error("Sorted posts should not differ when sorted in ascending manner")
			t.Fail()
		}
	}
}

// TestSortingPostsDesc sorting posts in descending manner.
func TestSortPostsDesc(t *testing.T) {
	conf, err := config.LoadConfig(testConfig)
	SetConfig(conf)
	if err != nil {
		t.Error("Error loading test conf: " + err.Error())
		t.Fail()
	}

	sortTypes := serverConfig.APISettings.AcceptableSortBy
	for _, sortType := range sortTypes {
		examplePosts := getExamplePosts()
		sorted := sortPosts(examplePosts, "desc", sortType)
		if !reflect.DeepEqual(sorted[0], getExamplePosts()[2]) {
			t.Error("Posts are not sorted in descending manner")
			t.Fail()
		}
		if !reflect.DeepEqual(sorted[1], getExamplePosts()[1]) {
			t.Error("Posts are not sorted in descending manner")
			t.Fail()
		}
		if !reflect.DeepEqual(sorted[2], getExamplePosts()[0]) {
			t.Error("Posts are not sorted in descending manner")
			t.Fail()
		}
	}
}

// TestRemoveDuplicates tests the removing of duplicate posts from a post array
func TestRemoveDuplicates(t *testing.T) {
	conf, err := config.LoadConfig(testConfig)
	SetConfig(conf)
	if err != nil {
		t.Error("Error loading test conf: " + err.Error())
		t.Fail()
	}

	examplePosts := getExamplePosts()
	// Duplicate the last post and add it to the end of the test posts
	examplePosts = append(examplePosts, examplePosts[2])

	removedDupes := removeDuplicatePosts(examplePosts)

	if !reflect.DeepEqual(removedDupes, getExamplePosts()) {
		t.Error("Duplicate post was not removed")
		t.Fail()
	}
}