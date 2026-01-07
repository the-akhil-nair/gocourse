package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type rateLimiterPseudo struct {
	ipCount map[string]uint16
	limit   uint16
}

func RateLimiterPseudo(next http.Handler) http.Handler {
	rl := rateLimiterPseudo{
		limit:   5,
		ipCount: make(map[string]uint16),
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		go func() {
			ticker := time.NewTicker(10 * time.Second)

			for range ticker.C {
				for k, _ := range rl.ipCount {
					rl.ipCount[k] = 0
				}
			}
		}()

		fmt.Println("Remote Address", r.RemoteAddr)
		if _, ok := rl.ipCount[r.RemoteAddr]; ok {
			if rl.ipCount[r.RemoteAddr] <= rl.limit {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "Rate Limit Reached", http.StatusForbidden)
				log.Println("Rate Limit Reached.")
				log.Printf("For Host%v\n", rl)
			}
			rl.ipCount[r.RemoteAddr]++
		} else {
			rl.ipCount[r.RemoteAddr] = 1

		}

	})
}
