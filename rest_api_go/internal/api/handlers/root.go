package handlers

import "net/http"

func RootHandler(w http.ResponseWriter, r *http.Request) {
	// _, err := fmt.Fprint(w, "Serving /\n")

	// if err != nil {
	// 	log.Fatalln("Unable to server /")
	// }
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello Get method on root route.\n"))
	case http.MethodPost:
		w.Write([]byte("Hello Post method on root route.\n"))
	case http.MethodPut:
		w.Write([]byte("Hello Put method on root route.\n"))
	case http.MethodPatch:
		w.Write([]byte("Hello Put method on root route.\n"))
	case http.MethodDelete:
		w.Write([]byte("Hello Delete method on root route.\n"))
	default:
		w.Write([]byte("Unknown HTTP Method is called."))
	}
}
