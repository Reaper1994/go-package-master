package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/treblle/treblle-go"
)

// LoggingMiddleware logs each incoming request.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// RecoveryMiddleware recovers from any panics and writes a 500 error.
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// TreblleMiddleware wraps the given handler with Treblle middleware for Api ops.
func TreblleMiddleware(apiKey string, projectId string, next http.Handler) http.Handler {

	treblle.Configure(treblle.Configuration{
		APIKey:    apiKey,
		ProjectID: projectId,
	})

	return treblle.Middleware(next)
}

const AuthToken = "GSLC-123-0R" //Dummy token set for now

// AuthorizationMiddleware checks if the request contains the correct token in the header.
func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		// Check if the Authorization header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Extract the token from the header
		token := strings.TrimPrefix(authHeader, "Bearer ")

		if token != AuthToken {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If the token is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
