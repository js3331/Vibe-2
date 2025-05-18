package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"vibe/internal/handlers"
)

func getEnvDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func main() {
	// Database connection
	dbHost := getEnvDefault("DB_HOST", "localhost")
	dbUser := getEnvDefault("DB_USER", "postgres")
	dbPassword := getEnvDefault("DB_PASSWORD", "postgres")
	dbName := getEnvDefault("DB_NAME", "vibe")
	dbPort := getEnvDefault("DB_PORT", "5432")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Wait for database to be ready
	for i := 0; i < 30; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		log.Printf("Waiting for database... %v", err)
		time.Sleep(time.Second)
	}
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Create router
	r := mux.NewRouter()

	// Setup handlers
	h := handlers.NewHandler(db)
	
	// API routes
	r.HandleFunc("/api/posts", h.CreatePost).Methods("POST")
	r.HandleFunc("/api/posts", h.SearchPosts).Methods("GET")

	// Serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.HandleFunc("/", h.HomePage).Methods("GET")

	// Start server
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
