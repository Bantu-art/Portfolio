package main

import (
	"log"
	"net/http"

	portfolio "portfolio/cv"
)
func main() {
	mux := http.NewServeMux()

	// Handle the homepage route
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			portfolio.NotFound(w, r)
			return
		}
		portfolio.HomePage(w, r)
	})

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Wrap mux with error handling middleware
	handler := portfolio.ErrorMiddleware(mux)

	// Start the server
	log.Println("Starting server on http://localhost:8080")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
