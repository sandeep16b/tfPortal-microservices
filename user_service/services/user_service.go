package services

import (
	"fmt"
	"goproject/user_service/data"
	"goproject/user_service/models"
)

// GetUsers returns all users
func GetAllUsers() ([]models.User, error) {
	fmt.Println("start: GetAllUsers")
	var users, err = data.GetAllUsers()
	fmt.Println("end: GetAllUsers")
	if err != nil {
		fmt.Println("Error fetching users: ", err)
		return nil, err
	}
	return users, err
}

// GetUser fetches a single user by ID
func GetUser(id int) (models.User, error) {
	return data.GetUser(id)
}

// CreateUser creates a new user
func CreateUser(user models.User) (string, error) {
	var err = data.CreateUser(user)
	if err != nil {
		fmt.Println("Error fetching users: ", err)
		return "", err
	}
	return "User created successfully.", err
}
