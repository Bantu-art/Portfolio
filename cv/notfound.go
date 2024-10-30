package portfolio

import "net/http"

// Custom NotFound handler for 404 errors
func NotFound(w http.ResponseWriter, r *http.Request) {
	ErrorHandler(w, r, http.StatusNotFound)
}

// Custom InternalServerError handler for 500 errors
func InternalServerError(w http.ResponseWriter, r *http.Request) {
	ErrorHandler(w, r, http.StatusInternalServerError)
}
