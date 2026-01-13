package middlewares

import (
	"log"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func ResponseTimer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received Response from ResponseTimer Middleware")
		start := time.Now()
		wrappedWriter := &responseWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(wrappedWriter, r)
		duration := time.Since(start)
		wrappedWriter.Header().Set("X-Response-Time", duration.String())
		log.Printf("Response Time: %s %s took %v with status %d\n", r.Method, r.URL.Path, duration, wrappedWriter.status)
		log.Println("Sent Response from ResponseTimer Middleware")
	})
}
