package handlers

import (
	"log"
	"net/http"
)

func StudentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Students route.\n"))
	log.Println("Hello, Students route.")

	// Handle different HTTP Methods

	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello Get method on Students route.\n"))
	case http.MethodPost:
		w.Write([]byte("Hello Post method on Students route.\n"))
	case http.MethodPut:
		w.Write([]byte("Hello Put method on Students route.\n"))
	case http.MethodPatch:
		w.Write([]byte("Hello Put method on Students route.\n"))
	case http.MethodDelete:
		w.Write([]byte("Hello Delete method on Students route.\n"))
	default:
		w.Write([]byte("Unknown HTTP Method is called."))
	}

}
