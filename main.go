package main

import (
	"fmt"
	"goproject/handlers"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Routing to /posts")
	http.HandleFunc("/posts", handlers.PostsHandler)
	fmt.Println("Server running at http://localhost:8092")
	log.Fatal(http.ListenAndServe(":8092", nil))
}
