package middlewares

import (
	"log"
	"net/http"
	"sync"
	"time"
)

type rateLimiter struct {
	mu        sync.Mutex
	visitors  map[string]int
	limit     int
	resetTime time.Duration
}

func NewRateLimiter(limit int, resetTime time.Duration) *rateLimiter {
	rl := &rateLimiter{
		limit:     limit,
		resetTime: resetTime,
		visitors:  make(map[string]int),
	}
	go rl.resetVisitorCount()
	return rl
}

func (rl *rateLimiter) resetVisitorCount() {
	for {
		time.Sleep(rl.resetTime)
		rl.mu.Lock()
		rl.visitors = make(map[string]int)
		rl.mu.Unlock()
	}
}

func (rl *rateLimiter) RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rl.mu.Lock()
		defer rl.mu.Unlock()

		visitorIP := r.RemoteAddr
		rl.visitors[visitorIP]++
		log.Printf("Visitor count for %v is %v", visitorIP, rl.visitors[visitorIP])

		if rl.visitors[visitorIP] > rl.limit {
			http.Error(w, "Too many Requsets", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
