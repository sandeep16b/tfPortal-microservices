package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

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
		log.Fatalf("Error opening DB connection: %v", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatalf("Cannot connect to MySQL: %v", err)
	}

	fmt.Println("✅ Connected to MySQL successfully")
}
