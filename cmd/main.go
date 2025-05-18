package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"vibe/internal/handlers"
)

func main() {
	// Database connection
	connStr := "postgres://postgres:postgres@localhost:5432/vibe?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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
