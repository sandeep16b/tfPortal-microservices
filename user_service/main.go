package main

import (
	"fmt"
	"goproject/user_service/data"
	"goproject/user_service/handlers"
	"path/filepath"
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

	loadEnv();
	
	fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))
	data.InitDB()
	mux := http.NewServeMux()

	fsUser := http.StripPrefix("/user/", http.FileServer(http.Dir("../public/user")))
	http.Handle("/user/", fsUser)

	fsJS := http.StripPrefix("/js/", http.FileServer(http.Dir("../public/js")))
	http.Handle("/js/", fsJS)
	mux.HandleFunc("/users", handlers.UsersHandler) // register API first

	log.Println("✅ Serving at http://localhost:8093")
	log.Fatal(http.ListenAndServe(":8093", nil))

}
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
