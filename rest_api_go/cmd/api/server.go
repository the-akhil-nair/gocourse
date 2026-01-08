package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	mw "restapi/internal/api/middlewares"
	"strings"
	"time"
)

type User struct {
	Name string `json:"name"`
	Age  uint16 `json:"age"`
	City string `json:"city"`
}

func utilityFunction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is a utility function.")
	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		//log.Fatalln("Unable to parse form data", err)
		http.Error(w, "Unable to parse form data", http.StatusBadRequest)
		return
	}

	response := make(map[string]interface{})

	for key, value := range r.Form {
		response[key] = value
	}

	fmt.Println("Response Data: ", response)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read body data", http.StatusBadRequest)
		return
	}

	fmt.Println("Body Data: ", string(body))

	decoder := json.NewDecoder(r.Body)
	var user User
	err = decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Unable to parse JSON body", http.StatusBadRequest)
		return
	}

	fmt.Println("User Data: ", user)
	fmt.Println("User Name: ", user.Name)
	fmt.Println("User Age: ", user.Age)
	fmt.Println("User City: ", user.City)

	response1 := make(map[string]interface{})

	err = json.Unmarshal(body, &response1)
	if err != nil {
		http.Error(w, "Unable to parse JSON body", http.StatusBadRequest)
		return
	}

	fmt.Println("Response1 Data: ", response1)

	// Parse the raw body data
	// body := r.Body
	// defer body.Close()

	// bodyData := make([]byte, r.ContentLength)
	// _, err = body.Read(bodyData)
	// if err != nil {
	// 	log.Fatalln("Unable to read body data", err)
	// }

	// fmt.Println("Body Data: ", string(bodyData))

	defer r.Body.Close()

	// Access the request details
	fmt.Println("Request Body: ", r.Body)
	fmt.Print("Requst form:", r.Form)
	fmt.Println("Request Content Length: ", r.ContentLength)
	fmt.Println("Request URL: ", r.URL)
	fmt.Println("Request Header: ", r.Header)
	fmt.Println("Request Context: ", r.Context())
	fmt.Println("Request Host: ", r.Host)
	fmt.Println("Request Method: ", r.Method)
	fmt.Println("Request Protocol: ", r.Proto)
	fmt.Println("Request Remote Address: ", r.RemoteAddr)
	fmt.Println("Request Request URI: ", r.RequestURI)
	fmt.Print("Requst TLS: ", r.TLS)
	fmt.Println("Request Transfer Encoding: ", r.TransferEncoding)
	fmt.Println("Request Trailers", r.Trailer)
	fmt.Println("Request User Agent", r.UserAgent())
	fmt.Println("Request Port: ", r.URL.Port())
	fmt.Println("Request URL Scheme: ", r.URL.Scheme)
}

func queryHandler(r *http.Request) {
	// teachers/{id}
	// teachers/9
	// teachers/?key=value&query=search&sortby=email&sortorder=asc
	// /query?key=value&query=search&sortby=email&sortorder=asc
	fmt.Println("Teachers Route Method: ", r.URL.Path)
	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	userID := strings.TrimSuffix(path, "/")
	queryParams := r.URL.Query()
	sortBy := queryParams.Get("sortby")
	sortOrder := queryParams.Get("sortorder")
	key := queryParams.Get("key")
	query := queryParams.Get("query")

	if sortOrder == "" {
		sortOrder = "asc"
	}

	fmt.Println("Query: ", query)
	fmt.Println("Key: ", key)
	fmt.Println("User ID: ", userID)
	fmt.Println("Sort By: ", sortBy)
	fmt.Println("Sort Order: ", sortOrder)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
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

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, teachers route.\n"))
	log.Println("Hello, teachers route.")

	switch r.Method {
	case http.MethodGet:
		// queryHandler(r)
		w.Write([]byte("Hello Get method on Teachers route.\n"))
	case http.MethodPost:
		w.Write([]byte("Hello Post method on Teachers route.\n"))
	case http.MethodPut:
		w.Write([]byte("Hello Put method on Teachers route.\n"))
	case http.MethodPatch:
		w.Write([]byte("Hello Put method on Teachers route.\n"))
	case http.MethodDelete:
		w.Write([]byte("Hello Delete method on Teachers route.\n"))
	default:
		w.Write([]byte("Unknown HTTP Method is called."))
	}
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
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

func main() {
	port := 3000

	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	mux.HandleFunc("/teachers/", teachersHandler)

	mux.HandleFunc("/students/", studentsHandler)
	mux.HandleFunc("/execs/", ExecsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	rl := mw.NewRateLimiter(5, time.Minute)
	hppOptions := mw.HPPOptions{
		CheckQuery:                  true,
		CheckBody:                   true,
		CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
		Whitelist:                   []string{"sortBy", "sortOrder", "name", "age", "class"},
	}

	fmt.Println("Server is running on Port: ", port)

	//multiplexer := mw.Hpp(hppOptions)(rl.RateLimiter(mw.Compression(mw.ResponseTimer(mw.SecurityHeaders(mw.Cors(mux))))))
	multiplexer := mw.Cors(rl.RateLimiter(mw.ResponseTimer(mw.SecurityHeaders(mw.Compression(mw.Hpp(hppOptions)())))))

	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		Handler:   multiplexer,
		TLSConfig: tlsConfig,
	}

	err := server.ListenAndServeTLS(cert, key)

	if err != nil {
		log.Fatalln("Unable to start the server", err)
	}
}
