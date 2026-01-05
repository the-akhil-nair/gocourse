package rest

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request came to / with method", r.Method)
		fmt.Fprintln(w, "Hello World!")
	})

	// Listening to request and serving responses at 3000
	const port string = ":3000"

	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatalln("Not able to listen and server on port 3000", err)
		return
	}
}
