package data

import "goproject/models"

func GetAllUsers() ([]models.User, error) {
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
			return nil, err
		}
		users = append(users, user)
	}
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
