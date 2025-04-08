package main

import (
	"fmt"
	"goproject/data"
	"goproject/handlers"
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

	err = godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}
	fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))
	data.InitDB()
	http.HandleFunc("/posts", handlers.PostsHandler)
	http.HandleFunc("/users", handlers.UsersHandler)
	http.HandleFunc("/comments", handlers.CommentsHandler)

	fmt.Println("fmt: Server running at http://localhost:8092")
	http.Handle("/", http.FileServer(http.Dir("./public")))
	// log.Fatal(http.ListenAndServe("0.0.0.0:8092", nil))
}
