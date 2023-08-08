package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request from %s", r.RemoteAddr)
	fmt.Fprintf(w, "Hello, Kubernetes!")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handler)

	// Start the server in a separate goroutine
	go func() {
		log.Printf("Server listening on port %s...", port)
		err := http.ListenAndServe(":"+port, nil)
		if err != nil {
			log.Fatal("Failed to start the server")
		}
	}()

	log.Printf("Click on the link http://localhost:%s", port)

	// Use a channel to prevent the main function from exiting immediately
	done := make(chan bool)
	<-done
}
