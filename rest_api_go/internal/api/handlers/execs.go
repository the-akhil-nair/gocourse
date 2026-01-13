package handlers

import "net/http"

func ExecsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello Get method on Executives route.\n"))
	case http.MethodPost:
		w.Write([]byte("Hello Post method on Executives route.\n"))
	case http.MethodPut:
		w.Write([]byte("Hello Put method on Executives route.\n"))
	case http.MethodPatch:
		w.Write([]byte("Hello Put method on Executives route.\n"))
	case http.MethodDelete:
		w.Write([]byte("Hello Delete method on Executives route.\n"))
	default:
		w.Write([]byte("Unknown HTTP Method is called."))
	}
}
