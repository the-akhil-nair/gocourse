package middlewares

import (
	"compress/gzip"
	"log"
	"net/http"
	"strings"
)

func Compression(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the client accept the gzip encoding
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
		}

		// Set the response Header
		w.Header().Set("Content-Encoding", "gzip")

		gz := gzip.NewWriter(w)
		gw := &gzipResponseWriter{ResponseWriter: w, Writer: gz}

		defer gz.Close()
		// Wrap the Response Writer
		next.ServeHTTP(gw, r)
		log.Println("Sent Response from Compression Middleware.")
	})
}

// gzipResponseWriter wraps http.ResponseWriter to write gzipped responses.

type gzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (g *gzipResponseWriter) Write(b []byte) (int, error) {
	return g.Writer.Write(b)
}
