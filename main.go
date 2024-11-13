package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	portfolio "portfolio/cv"
)

func main() {
	iport := port()
	port := ":" + strconv.Itoa(iport)
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
	err := http.ListenAndServe(port, handler)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}

func port() int {
	defaultPort := 9000
	portStr := os.Getenv("PORT")
	var port int

	if portStr != "" {
		var err error
		port, err = strconv.Atoi(portStr)
		if err != nil {
			port = defaultPort // Use default if conversion fails
		}
	} else {
		port = defaultPort
	}
	return port
}
