package http

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

func loadClientCAs() *x509.CertPool {
	clientCAs := x509.NewCertPool()
	caCert, err := os.ReadFile("cert.pem")
	if err != nil {
		log.Fatalln("Could not load client CA:", err)
	}
	clientCAs.AppendCertsFromPEM(caCert)
	return clientCAs
}

func main() {

	// This will handle Any method as we have not explicitly mentioned the method.
	// Handle function should be above listenAndServe else we will get 404.
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		logRequestDetails(r)
		fmt.Fprintln(w, "Handling Incoming Orders.")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		logRequestDetails(r)

		fmt.Fprintln(w, "Handling Incoming Users.")
	})

	port := 3000

	// Load the TLS cert and key
	cert := "cert.pem"
	key := "key.pem"

	// Configure the TLS
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		ClientCAs:  loadClientCAs(),
		ClientAuth: tls.RequireAndVerifyClientCert,
	}

	// Create the custom server
	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		Handler:   nil,
		TLSConfig: tlsConfig,
	}

	// Enable http2
	// ALPN : Application Layer protocol Negotiaiotion.
	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Server is running on Port: ", port)

	err := server.ListenAndServeTLS(cert, key)

	if err != nil {
		log.Fatalln("Unable to start the HTTP2 sever with TLS.")
		return
	}
	// HTTP1.1 server without TLS.
	// err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	// if err != nil {
	// 	log.Fatalln("Not able to listen and server on port 3000", err)
	// 	return
	// }
}

func logRequestDetails(r *http.Request) {
	httpVersion := r.Proto
	log.Println("Recieved the Request with HTTP Version", httpVersion)

	if r.TLS != nil {
		tlsVersion := getTLSVersionName(r.TLS.Version)
		log.Println(tlsVersion)
	} else {
		log.Println("Recieved without TLS.")
	}

}

func getTLSVersionName(version uint16) string {
	switch version {
	case tls.VersionTLS10:
		return "Version 1.0"
	case tls.VersionTLS11:
		return "Version 1.1"
	case tls.VersionTLS12:
		return "Version 1.2"
	case tls.VersionTLS13:
		return "Version 1.3"
	default:
		return "Unknown TLS version"
	}
}
