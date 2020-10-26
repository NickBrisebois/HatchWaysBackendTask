// Package handlers provides handlers for different API routes
package handlers

import (
	"context"
	"encoding/json"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"net/http"
	"sort"
)

// Post is a struct representing a single JSON Post data object
type Post struct {
	// Author is the author of the blog post
	Author     string   `json:"author"`
	// AuthorID is the author ID of the post
	AuthorID   int      `json:"authorId"`
	// ID is the ID of the post
	ID         int      `json:"id"`
	// Likes represents how many people have liked this post
	Likes      int      `json:"likes"`
	// Popularity is a rating of how popular this post is
	Popularity float64  `json:"popularity"`
	// Reads represents how many people have read this post
	Reads      int      `json:"reads"`
	// Tags are tags labelling this post
	Tags       []string `json:"tags"`
}

// Posts is a struct that acts as a container for an array of Posts data
type Posts struct {
	// Posts holds a bunch of posts.
	Posts []Post `json:"posts"`
}

// getPosts gets all posts from given tag from HatchWays API
func getPosts(tag string) (*Posts, error) {
	// URL To make GET request to with tag as query
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

	// Convert posts JSON to struct data
	var posts Posts
	if err = json.Unmarshal(rawPosts, &posts); err != nil {
		return nil, err
	}

	return &posts, nil
}

// getCombinedPosts retrieves all posts from all of the given tags
func getCombinedPosts(tags []string) ([]Post, error) {
	var combinedPosts []Post

	// Create an error group so that we can use goroutines to concurrently
	// fetch posts and any errors that occur
	apiCallErrors, _ := errgroup.WithContext(context.Background())
	// Loop through all given tags and combine them
	for _, tag := range tags {
		// Spawn new goroutine to retrieve posts under this tag
		apiCallErrors.Go(func() error {
			// Get posts from this tag
			postsFromTag, err := getPosts(tag)
			if err != nil {
				return err
			}

			// Add all posts from this tag to combined list of posts
			combinedPosts = append(combinedPosts, postsFromTag.Posts...)
			return nil
		})
	}
	// Wait for goroutines to finish then check for errors
	err := apiCallErrors.Wait()
	if err != nil {
		// This will be the first (or only) error that occurred when retrieving posts
		return nil, err
	}

	// Remove dupes
	postsNoDuplicates := removeDuplicatePosts(combinedPosts)
	return postsNoDuplicates, nil
}

// removeDuplicatePosts removes duplicates from given posts using post ID
func removeDuplicatePosts(posts []Post) []Post {
	occurredPosts := map[int]bool{}
	var result []Post

	for _, post := range posts {
		if !occurredPosts[post.ID] {
			occurredPosts[post.ID] = true
			result = append(result, post)
		}
	}

	return result
}

// sortPosts sorts given posts in ascending or descending manner by the selected sortBy value
func sortPosts(posts []Post, direction string, sortBy string) []Post {
	switch sortBy {
	// Sort using built in Go sorting interface
	// https://golang.org/pkg/sort
	case "id":
		sort.Slice(posts, func(i, j int) bool {
			// Anonymous function implements Less function in sorting interface
			// https://golang.org/pkg/sort/#StringSlice.Less
			if direction == "asc" {
				return posts[i].ID < posts[j].ID
			}else {
				return posts[i].ID > posts[j].ID
			}
		})
	case "reads":
		sort.Slice(posts, func(i, j int) bool {
			if direction == "asc" {
				return posts[i].Reads < posts[j].Reads
			}else {
				return posts[i].Reads > posts[j].Reads
			}
		})
	case "likes":
		sort.Slice(posts, func(i, j int) bool {
			if direction == "asc" {
				return posts[i].Likes < posts[j].Likes
			}else {
				return posts[i].Likes > posts[j].Likes
			}
		})
	case "popularity":
		sort.Slice(posts, func(i, j int) bool {
			if direction == "asc" {
				return posts[i].Popularity < posts[j].Popularity
			}else {
				return posts[i].Popularity > posts[j].Popularity
			}
		})
	}

	return posts
}
