package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Tutorial struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Set up routes
	http.HandleFunc("/api/health", corsMiddleware(healthHandler))
	http.HandleFunc("/api/tutorials", corsMiddleware(tutorialsHandler))

	log.Printf("Server starting on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

// corsMiddleware adds CORS headers to allow frontend access
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get allowed origin from environment variable, default to localhost for development
		allowedOrigin := os.Getenv("ALLOWED_ORIGIN")
		if allowedOrigin == "" {
			allowedOrigin = "*"
		}
		
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

// healthHandler returns the health status of the API
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: "API is running",
	}); err != nil {
		log.Printf("Error encoding health response: %v", err)
	}
}

// tutorialsHandler returns a list of tutorials
func tutorialsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Message: "Method not allowed",
		})
		return
	}

	// Sample tutorials data
	tutorials := []Tutorial{
		{
			ID:          1,
			Title:       "Getting Started with Go",
			Description: "Learn the basics of Go programming language",
			URL:         "https://go.dev/doc/tutorial/getting-started",
		},
		{
			ID:          2,
			Title:       "React Fundamentals",
			Description: "Master the core concepts of React",
			URL:         "https://react.dev/learn",
		},
		{
			ID:          3,
			Title:       "Building REST APIs",
			Description: "Create robust REST APIs with Go",
			URL:         "https://go.dev/doc/tutorial/web-service-gin",
		},
		{
			ID:          4,
			Title:       "Modern JavaScript",
			Description: "ES6+ features and best practices",
			URL:         "https://javascript.info",
		},
	}

	if err := json.NewEncoder(w).Encode(Response{
		Success: true,
		Data:    tutorials,
	}); err != nil {
		log.Printf("Error encoding tutorials response: %v", err)
	}
}
