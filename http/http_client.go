package rest

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Create a new http client

	client := &http.Client{}

	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts/1")

	// resp, err := client.Get("https://swapi.dev/api/people/1")

	if err != nil {
		fmt.Println("Error making get request", err)
		return
	}

	defer resp.Body.Close()

	// Read and print the response body
	text, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error in reading the response", err)
		return
	}

	fmt.Println(string(text))
}
