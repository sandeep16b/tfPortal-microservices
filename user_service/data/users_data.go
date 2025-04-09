package data

import (
	"fmt"
	"goproject/user_service/models"
)

func GetAllUsers() ([]models.User, error) {
    fmt.Println("Running query to get users...")  // Debug line

    rows, err := DB.Query("SELECT id, name, email FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        err := rows.Scan(&user.ID, &user.Name, &user.Email)
        if err != nil {
			fmt.Println("Error executing query:", err)  // Debug line
            return nil, err
        }
        users = append(users, user)
    }

    fmt.Println("Fetched users: ", users)  // Debug line
    return users, nil
}

// GetUser retrieves a user by ID from the database
func GetUser(id int) (models.User, error) {
	var user models.User
	row := DB.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

// CreateUser adds a new user to the database
func CreateUser(user models.User) error {
	_, err := DB.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	return err
}
