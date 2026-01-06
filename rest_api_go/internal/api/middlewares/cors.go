package middlewares

import (
	"fmt"
	"net/http"
)

var allowedOrigins = []string{
	"https://localhost:3000",
	"http://example.com",
	"https://my-origin-url.com",
}

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		fmt.Println("CORS Middleware Origin:", origin)

		if isOriginAllowed(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			http.Error(w, "CORS origin denied", http.StatusForbidden)
			return
		}
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		w.Header().Set("Max-Age", "3600")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func isOriginAllowed(origin string) bool {
	for _, allowedOrigin := range allowedOrigins {
		if origin == allowedOrigin {
			return true
		}
	}
	return false
}
