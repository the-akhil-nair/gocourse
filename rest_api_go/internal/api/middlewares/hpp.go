package middlewares

import (
	"log"
	"net/http"
	"strings"
)

type HPPOptions struct {
	CheckQuery                  bool
	CheckBody                   bool
	CheckBodyOnlyForContentType string
	Whitelist                   []string
}

func Hpp(options HPPOptions) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if options.CheckBody && r.Method == http.MethodPost && isCorrectContentType(r, options.CheckBodyOnlyForContentType) {
				// Filter the body parameters.
				filterBodyParams(r, options.Whitelist)
			}
			if options.CheckQuery && r.URL.Query() != nil {
				// Filter the query parameters.
				filterQueryParams(r, options.Whitelist)
			}
			next.ServeHTTP(w, r)
		})
	}
}

func isCorrectContentType(r *http.Request, contentType string) bool {
	return strings.Contains(r.Header.Get("Content-Type"), contentType)
}

func filterBodyParams(r *http.Request, whitelist []string) {
	err := r.ParseForm()

	if err != nil {
		log.Println(err)
		return
	}

	for k, v := range r.Form {
		if len(v) > 1 {
			r.Form.Set(k, v[0]) // First value.
			// r.Form.Set(k, v[len(v)-1]) // Last value.
		}
		if !isWhitelisted(k, whitelist) {
			delete(r.Form, k)
		}
	}
}

func filterQueryParams(r *http.Request, whitelist []string) {
	query := r.URL.Query()

	for k, v := range query {
		if len(v) > 1 {
			query.Set(k, v[0]) // First value.
			// r.Form.Set(k, v[len(v)-1]) // Last value.
		}
		if !isWhitelisted(k, whitelist) {
			query.Del(k)
		}
	}
	r.URL.RawQuery = query.Encode()
}

func isWhitelisted(param string, whitelist []string) bool {
	for _, v := range whitelist {
		if param == v {
			return true
		}
	}
	return false
}
