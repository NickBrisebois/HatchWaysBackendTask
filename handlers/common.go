package handlers

import (
	"encoding/json"
	_ "encoding/json"
	"github.com/NickBrisebois/HatchWaysAppBackend/config"
	"io/ioutil"
	"net/http"
)

type PostsRetriever struct {
	config *config.Config
}

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

func NewPostsRetriever(config *config.Config) *PostsRetriever {
	return &PostsRetriever{
		config: config,
	}
}

func (pr *PostsRetriever) GetPosts(tag string) (*Posts, error) {
	getPostsURL := pr.config.Incoming.DataSrc + "?tag=" + tag
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
