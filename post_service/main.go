package main

import (
	"fmt"
	"goproject/post_service/handlers"
	"path/filepath"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)
func loadEnv() {
	// Dynamically walk up from current dir to find `.env`
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("❌ Failed to get working directory: %v", err)
	}

	for {
		envPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			err = godotenv.Load(envPath)
			if err != nil {
				log.Fatalf("❌ Error loading .env: %v", err)
			}
			fmt.Println("✅ Loaded .env from:", envPath)
			return
		}

		// Reached root and didn't find it
		parent := filepath.Dir(dir)
		if parent == dir {
			log.Fatal("❌ .env file not found in any parent directory")
		}
		dir = parent
	}
}

func main() {
	fmt.Println("Routing to /posts")
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current directory:", dir)

	loadEnv();
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}
	fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))
	// data.InitDB()
	mux := http.NewServeMux()

	fsPost := http.StripPrefix("/post/", http.FileServer(http.Dir("../public/post")))
	http.Handle("/post/", fsPost)

	fsJS := http.StripPrefix("/js/", http.FileServer(http.Dir("../public/js")))
	http.Handle("/js/", fsJS)
	mux.HandleFunc("/posts", handlers.PostsHandler) // register API first

	log.Println("✅ Serving at http://localhost:8092")
	log.Fatal(http.ListenAndServe(":8092", nil))
}
