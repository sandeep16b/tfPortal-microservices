package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var DB *sql.DB

func InitDB() {
	var err error
	connString := "server=localhost\\SQLEXPRESS;database=GoPostDB;Trusted_Connection=Yes;encrypt=disable"

	DB, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("❌ Error creating connection pool: ", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("❌ Cannot connect to SQL Server: ", err)
	}

	fmt.Println("✅ Connected to SQL Server Express!")
}
