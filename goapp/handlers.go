package main

import (
	"fmt"
	"html"
	"net/http"
)

// RootHandler handles requests to the root endpoint
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// HiHandler handles requests to the /hi endpoint
func HiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi There Yolanda!")
}