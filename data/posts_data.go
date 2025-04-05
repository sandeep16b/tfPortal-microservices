package data

import (
	"bytes"
	"encoding/json"
	"net/http"

	"goproject/models"
)

func FetchPostsFromAPI() ([]models.Post, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var posts []models.Post
	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		return nil, err
	}
	return posts, nil
}

func SavePostToAPI(post models.Post) (models.Post, error) {
	data, err := json.Marshal(post)
	if err != nil {
		return models.Post{}, err
	}

	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return models.Post{}, err
	}
	defer resp.Body.Close()

	var created models.Post
	if err := json.NewDecoder(resp.Body).Decode(&created); err != nil {
		return models.Post{}, err
	}
	return created, nil
}
