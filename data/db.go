package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Initialize the database connection
func InitDB() {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbname := os.Getenv("MYSQL_DB")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbname)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	// If you have to create or update schema in db, uncomment below line
	InitTables()

	// Ensure the DB is reachable
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error pinging the database: ", err)
	}

	fmt.Println("✅ Connected to MySQL successfully")
}

// Create the necessary tables if they do not exist
func InitTables() {

	// Creating Comments table
	_, err := DB.Exec(`
        CREATE TABLE IF NOT EXISTS comments (
            id INT AUTO_INCREMENT PRIMARY KEY,
            comment VARCHAR(300) NOT NULL,
            postId int NOT NULL UNIQUE
        );
    `)
	if err != nil {
		log.Fatal("Error creating comments table: ", err)
	}

	// Creating Users table
	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS users (
            id INT AUTO_INCREMENT PRIMARY KEY,
            name VARCHAR(100) NOT NULL,
            email VARCHAR(100) NOT NULL UNIQUE
        );
	`)
	if err != nil {
		log.Fatal("Error creating posts table: ", err)
	}

	// Creating Posts table (if needed)
	_, err = DB.Exec(`
        CREATE TABLE IF NOT EXISTS posts (
            id INT AUTO_INCREMENT PRIMARY KEY,
            title VARCHAR(100),
            content TEXT,
            user_id INT,
            FOREIGN KEY (user_id) REFERENCES users(id)
        );
    `)
	if err != nil {
		log.Fatal("Error creating posts table: ", err)
	}
}
