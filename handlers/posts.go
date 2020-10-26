package handlers

import (
	"encoding/json"
	_ "encoding/json"
	"io/ioutil"
	"net/http"
)

type Post struct {
	Author     string   `json:"author"`
	AuthorID   int      `json:"authorId"`
	ID         int      `json:"id"`
	Likes      int      `json:"likes"`
	Popularity float64  `json:"popularity"`
	Reads      int      `json:"reads"`
	Tags       []string `json:"tags"`
}

type Posts struct {
	Posts []Post `json:"posts"`
}

// getPosts gets all posts from given tag from HatchWays API
func getPosts(tag string) (*Posts, error) {
	getPostsURL := serverConfig.Incoming.DataSrc + "?tag=" + tag
	resp, err := http.Get(getPostsURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	rawPosts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var posts Posts
	if err = json.Unmarshal(rawPosts, &posts); err != nil {
		return nil, err
	}

	return &posts, nil
}

// getCombinedPosts retrieves all posts from all of the given tags
func getCombinedPosts(tags []string) ([]Post, error) {
	var combinedPosts []Post
	// Loop through all given tags and combine them
	for _, tag := range tags {
		// Get posts from this tag
		postsFromTag, err := getPosts(tag)
		if err != nil {
			return nil, err
		}

		// Add all posts from this tag to combined list of posts
		combinedPosts = append(combinedPosts, postsFromTag.Posts...)
	}

	return combinedPosts, nil
}
