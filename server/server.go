package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Setup handler to listen the root path
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("request passed")
		fmt.Fprint(writer, "Hello from server")
	})

	// Server on port 9090 of local host
	server := http.Server{
		Addr:    "localhost:9091",
		Handler: handler,
	}
	if err := server.ListenAndServeTLS("/home/appscodepc/go/src/github.com/sayedppqq/mTLS/certs/server.crt",
		"/home/appscodepc/go/src/github.com/sayedppqq/mTLS/certs/server.key"); err != nil {
		log.Fatalf("error listening from server: %v", err)
	}
}
