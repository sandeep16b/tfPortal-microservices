package models

// Comment model
type Comment struct {
	ID      int    `json:"id"`
	Comment string `json:"comment"`
	PostId  int    `json:"postId"`
}
