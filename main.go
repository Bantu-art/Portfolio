package main

import (
	"log"
	"net/http"

	portfolio "portfolio/cv"
)

func main() {
	// Handle the homepage route
	http.HandleFunc("/", portfolio.HomePage)

	// Serve static files like CSS, JavaScript, and images
	fs := http.FileServer(http.Dir("static")) // Folder where static files are located
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Start the server
	log.Println("Starting server on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
