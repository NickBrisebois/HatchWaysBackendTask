package main

import (
	"fmt"
	"github.com/NickBrisebois/HatchWaysAppBackend/config"
	"io/ioutil"
	"net/http"
)

type PostsRetriever struct {
	config *config.Config
}

type Post struct {

}

func NewPostsRetriever(config *config.Config) *PostsRetriever {
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

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Print(string(body))

	return nil
}
