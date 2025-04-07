package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"goproject/models"
	"goproject/services"
)

// UsersHandler handles the user-related routes
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request is for a specific user by ID or for all users
	switch r.Method {
	case http.MethodGet:
		// Check if an ID parameter is provided in the URL path
		id := r.URL.Path[len("/users/"):]

		// If ID is provided, retrieve a single user
		if id != "" {
			userID, err := strconv.Atoi(id)
			if err != nil {
				http.Error(w, "Invalid user ID", http.StatusBadRequest)
				return
			}

			user, err := services.GetUser(userID) // Fetch user by ID from service layer
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
		} else {
			// Otherwise, fetch all users
			users, err := services.GetAllUsers() // Fetch all users from service layer
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(users)
		}

	case http.MethodPost:
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		created, err := services.CreateUser(user) // Save the new user to the database
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
