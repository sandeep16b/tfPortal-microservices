package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"goproject/models"
	"goproject/services"
)

// CommentHandler handles the comment-related routes
func CommentsHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request is for a specific comment by ID or for all comments
	switch r.Method {
	case http.MethodGet:
		fmt.Println("reached comment get")
		// Check if an ID parameter is provided in the URL path
		id := r.URL.Path[len("/comments/"):]

		// If ID is provided, retrieve a single comment
		if id != "" {
			postID, err := strconv.Atoi(id)
			if err != nil {
				http.Error(w, "Invalid comment ID", http.StatusBadRequest)
				return
			}

			comment, err := services.GetComment(postID) // Fetch comment by ID from service layer
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(comment)
		} else {
			// Otherwise, fetch all comments
			comments, err := services.GetAllComments() // Fetch all comments from service layer
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(comments)
		}

	case http.MethodPost:
		var comment models.Comment
		if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		created, err := services.CreateComment(comment) // Save the new comment to the database
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(created)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
