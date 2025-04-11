package main

import (
	"fmt"
	"goproject/user_service/data"
	"goproject/user_service/handlers"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Routing to /posts")
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current directory:", dir)

	err = godotenv.Load("../.env")
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}
	fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))
	data.InitDB()
	http.HandleFunc("/users", handlers.UsersHandler)
	fmt.Println("fmt: Server running at http://localhost:8093")
	http.Handle("/", http.FileServer(http.Dir("../public")))
	log.Fatal(http.ListenAndServe("0.0.0.0:8093", nil))
}
