package main

import (
	"log"
	"net/http"
)

func main() {
	// Register handlers
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/hi", HiHandler)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}