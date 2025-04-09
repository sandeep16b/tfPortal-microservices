package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"goproject/user_service/models"
	"goproject/user_service/services"
)

// UsersHandler handles the user-related routes
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request is for a specific user by ID or for all users
	switch r.Method {
	case http.MethodGet:
		fmt.Println("r.URL.Path>>", r.URL.Path)
	
		// Check if the path length is long enough to contain the user ID after "/users"
		if len(r.URL.Path) <= len("/users") {
			// If no ID is present, fetch all users
			users, err := services.GetAllUsers() // Fetch all users from service layer
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
	
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(users)
			return
		}
	
		// Remove trailing slash if exists
		path := r.URL.Path
		if path[len(path)-1] == '/' {
			path = path[:len(path)-1]
		}
	
		// Get the ID after "/users/"
		id := path[len("/users/"):]
		fmt.Println("id>>", id)
	
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
