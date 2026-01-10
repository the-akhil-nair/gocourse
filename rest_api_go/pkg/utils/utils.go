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