package data

import (
	"goproject/models"
)

func FetchPostsFromDB() ([]models.Post, error) {
	rows, err := DB.Query("SELECT id, userId, title, body FROM Posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Body)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func SavePostToDB(post models.Post) (models.Post, error) {
	query := "INSERT INTO Posts (userId, title, body) OUTPUT INSERTED.id VALUES (@p1, @p2, @p3)"
	err := DB.QueryRow(query, post.UserID, post.Title, post.Body).Scan(&post.ID)
	if err != nil {
		return models.Post{}, err
	}
	return post, nil
}
