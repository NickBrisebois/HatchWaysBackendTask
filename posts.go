package main

import (
	"go/types"
	"net/http"
)

type PostsRetriever struct {
	config *types.Config
}

type Post struct {

}

func NewPostsRetriever(config *types.Config) *PostsRetriever {
	return &PostsRetriever{
		config: config,
	}
}

func (pr *PostsRetriever) GetPosts() error {
	resp, err := http.Get(pr.config.Incoming.DataSrc)
	_ = resp
	if err != nil {
		return err
	}

	return nil
}
