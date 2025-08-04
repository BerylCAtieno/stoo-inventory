package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")
	log.Printf("Received request from %s for path: %s", r.RemoteAddr, r.URL.Path)
	fmt.Fprintf(w, "Welcome to Stoo Inventory Management App")
}

func main() {
	http.HandleFunc("/", helloHandler)

	log.Println("Starting server on port 8080...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
