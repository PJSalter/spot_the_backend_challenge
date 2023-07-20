// main.go

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

func main() {

	http.HandleFunc("/spots", getSpotsInArea)
	log.Fatal(http.ListenAndServe(":8080", nil))

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

// Access the database URL from the environment variable
dbURL := os.Getenv("DATABASE_URL")

// Your database connection code here, using the dbURL
fmt.Println("Database URL:", dbURL)
}
