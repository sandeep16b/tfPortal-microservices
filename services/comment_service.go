package services

import (
	"fmt"
	"goproject/data"
	"goproject/models"
)

// GetComments returns all comments
func GetAllComments() ([]models.Comment, error) {
	var comments, err = data.GetAllComments()
	if err != nil {
		fmt.Println("Error fetching comments: ", err)
		return nil, err
	}
	return comments, err
}

// GetComment fetches a single comment by ID
func GetComment(id int) (models.Comment, error) {
	return data.GetComment(id)
}

// CreateComment creates a new comment
func CreateComment(comment models.Comment) (string, error) {
	var err = data.CreateComment(comment)
	if err != nil {
		fmt.Println("Error fetching comments: ", err)
		return "", err
	}
	return "Comment created successfully.", err
}
