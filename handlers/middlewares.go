package handlers

import (
	"fmt"
	"net/http"
)


const (
	DEFAULT_TOKEN = "token12345"
	DEFAULT_API_KEY = "key12345"
)


// Token Based Authentication
func TokenAuthMiddleware(next http.Handler) http.Handler {
	validToken := DEFAULT_TOKEN
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthoried. Token missing")
			return
		}
		if token != validToken {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized. Invalid Token")
			return
		}
		next.ServeHTTP(w, r)
	})
}


// API KEY Authentication
func ApiKeyAuthMiddleware(next http.Handler) http.Handler {
	validApiKey := DEFAULT_API_KEY
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		if apiKey == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthoried. API Key missing")
			return
		}
		if apiKey != validApiKey {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized. Invalid API Key")
			return
		}
		next.ServeHTTP(w, r)
	})
}


// JWT Authentication
func JwtAuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}


func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)  {
		next.ServeHTTP(w, r)
	})
}
