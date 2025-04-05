package main

import (
	"fmt"
	"goproject/data"
	"goproject/handlers"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Routing to /posts")
	data.InitDB()
	http.HandleFunc("/posts", handlers.PostsHandler)
	fmt.Println("Server running at http://localhost:8092")
	log.Fatal(http.ListenAndServe(":8092", nil))
}
