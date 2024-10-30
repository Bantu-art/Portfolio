package portfolio

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func TemplateError(w http.ResponseWriter, r *http.Request, templateName string) {
	log.Printf("Template error: %s is missing or inaccessible", templateName)
	data := ErrorPageData{
		StatusCode: http.StatusInternalServerError,
		StatusText: "Template Error",
		Message:    fmt.Sprintf("The template '%s' is missing or inaccessible. Our team has been notified.", templateName),
		RequestURL: r.URL.Path,
	}

	// Try to use error template
	if tmpl, err := template.ParseFiles("templates/error.html"); err == nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.Execute(w, data)
		return
	}

	// Fallback plain text response if error template is also missing
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Error 500: Template Error\n%s", data.Message)
}
