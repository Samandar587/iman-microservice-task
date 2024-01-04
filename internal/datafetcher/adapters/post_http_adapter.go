package adapters

import (
	"encoding/json"
	"golang-project-template/internal/datafetcher/domain"
	"io/ioutil"
	"log"
	"net/http"
)

type Meta struct {
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Page int `json:"page"`
}

type Post struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Response struct {
	Meta Meta   `json:"meta"`
	Data []Post `json:"data"`
}

type httpClient struct {
	f domain.PostFactory
}

func NewPostProvider() httpClient {
	return httpClient{}
}

func (h httpClient) GetPosts(page string) ([]domain.Post, error) {
	res, err := http.Get("https://gorest.co.in/public/v1/posts?page=" + page)
	if err != nil {
		return nil, err
	}
	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var response Response
	var postList []domain.Post
	json.Unmarshal(resData, &response)
	for _, post := range response.Data {
		domainPost := h.f.ParseToDomain(post.UserID, post.Title, post.Body, response.Meta.Pagination.Page)
		postList = append(postList, *domainPost)
	}
	return postList, nil
}
