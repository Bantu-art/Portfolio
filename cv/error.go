package portfolio

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// ErrorPageData holds the data for the error page template
type ErrorPageData struct {
	StatusCode int
	StatusText string
	Message    string
	RequestURL string // Added to show which URL caused the 404
}

// ErrorHandler handles both 404 and 500 errors
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	data := ErrorPageData{
		StatusCode: status,
		StatusText: http.StatusText(status),
		RequestURL: r.URL.Path,
	}

	switch status {
	case http.StatusNotFound:
		data.Message = fmt.Sprintf("The page '%s' doesn't exist. Please check the URL and try again.", r.URL.Path)
	case http.StatusInternalServerError:
		data.Message = "A template or resource file is missing. Our team has been notified."
		// Log the error for debugging
		log.Printf("500 error occurred: URL: %s, Error: Template missing or inaccessible", r.URL.Path)
	default:
		data.Message = "An unexpected error occurred. Please try again later."
	}

	// Check if error template exists before trying to use it
	if _, err := os.Stat("templates/error.html"); os.IsNotExist(err) {
		// Fallback response if error template is missing
		w.WriteHeader(status)
		fmt.Fprintf(w, "Error %d: %s\n%s", status, http.StatusText(status), data.Message)
		return
	}

	// Parse and execute error template
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		// Fallback response if template parsing fails
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error %d: %s\n%s", status, http.StatusText(status), data.Message)
		return
	}

	w.WriteHeader(status)
	tmpl.Execute(w, data)
}
