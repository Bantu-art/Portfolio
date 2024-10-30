package portfolio

import (
	"log"
	"net/http"
)

// Middleware to handle panics and errors
func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				InternalServerError(w, r)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
