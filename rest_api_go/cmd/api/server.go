package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 3000

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// _, err := fmt.Fprint(w, "Serving /\n")

		// if err != nil {
		// 	log.Fatalln("Unable to server /")
		// }

		w.Write([]byte("Hello, root route.\n"))
		log.Println("Hello, root route.")
	})

	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, teachers route.\n"))
		log.Println("Hello, teachers route.")
	})

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Students route.\n"))
		log.Println("Hello, Students route.")
		fmt.Println(r.Method)

		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("Hello Get method on Teachers route.\n"))
			log.Println("Get Method is called.")
		case http.MethodPost:
			w.Write([]byte("Hello Post method on Teachers route.\n"))
			log.Println("Post Method is called.")
		case http.MethodPut:
			w.Write([]byte("Hello Put method on Teachers route.\n"))
			log.Println("Put Method is called.")
		case http.MethodPatch:
			w.Write([]byte("Hello Put method on Teachers route.\n"))
			log.Println("Put Method is called.")
		default:
			w.Write([]byte("Unknown HTTP Method is called."))
			log.Fatalln("Unknown HTTP Method is called.")

		}

	})

	http.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Executives route.\n"))
		log.Println("Hello, Executives route.")
	})

	fmt.Println("Server is running on Port: ", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		log.Fatalln("Unable to start the server", err)
	}
}
