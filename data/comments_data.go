package data

import "goproject/models"

func GetAllComments() ([]models.Comment, error) {
	rows, err := DB.Query("SELECT id, comment, postId FROM comments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.ID, &comment.Comment, &comment.PostId)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

// GetComments retrieves a comment by ID from the database
func GetComment(id int) (models.Comment, error) {
	var comment models.Comment
	row := DB.QueryRow("SELECT id, comment, postId FROM comments WHERE id = ?", id)
	err := row.Scan(&comment.ID, &comment.Comment, &comment.PostId)
	if err != nil {
		return comment, err
	}
	return comment, nil
}

// CreateUser adds a new comment to the database
func CreateComment(comment models.Comment) error {
	_, err := DB.Exec("INSERT INTO comments (comment, postId) VALUES (?, ?)", comment.Comment, comment.PostId)
	return err
}
