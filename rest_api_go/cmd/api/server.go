package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	mw "restapi/internal/api/middlewares"
	"restapi/internal/repository/sqlconnect"
	"restapi/internal/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error in loading .env file", err)
	}

	db, err := sqlconnect.ConnectDB()

	if err != nil {
		panic(fmt.Sprintf("Unable to open Databse - %e", err))
	}

	log.Println(db)

	cert := "cert.pem"
	key := "key.pem"

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// rl := mw.NewRateLimiter(5, time.Minute)
	// hppOptions := mw.HPPOptions{
	// 	CheckQuery:                  true,
	// 	CheckBody:                   true,
	// 	CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
	// 	Whitelist:                   []string{"sortBy", "sortOrder", "name", "age", "class"},
	// }

	fmt.Println("Server is running on Port ", os.Getenv("API_PORT"))

	//multiplexer := mw.Hpp(hppOptions)(rl.RateLimiter(mw.Compression(mw.ResponseTimer(mw.SecurityHeaders(mw.Cors(mux))))))
	//multiplexer := mw.Cors(rl.RateLimiter(mw.ResponseTimer(mw.SecurityHeaders(mw.Compression(mw.Hpp(hppOptions)())))))
	// multiplexer := utils.ApplyMiddlewares(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeaders, mw.ResponseTimer, rl.RateLimiter, mw.Cors)
	multiplexer := mw.SecurityHeaders(router.Router())

	server := &http.Server{
		Addr:      os.Getenv("API_PORTl"),
		Handler:   multiplexer,
		TLSConfig: tlsConfig,
	}

	err = server.ListenAndServeTLS(cert, key)

	if err != nil {
		log.Fatalln("Unable to start the server", err)
	}
}
