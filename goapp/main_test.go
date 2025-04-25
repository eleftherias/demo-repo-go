package main

import (
	"fmt"
	"html"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootEndpoint(t *testing.T) {
	// Create a request to the root endpoint
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()
	
	// Create the handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	// Serve the request
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	expected := `Hello, "/"`
	body, _ := io.ReadAll(rr.Body)
	if string(body) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(body), expected)
	}
}

func TestHiEndpoint(t *testing.T) {
	// Create a request to the /hi endpoint
	req, err := http.NewRequest("GET", "/hi", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()
	
	// Create the handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi There Yolanda!")
	})

	// Serve the request
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	expected := "Hi There Yolanda!"
	body, _ := io.ReadAll(rr.Body)
	if string(body) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(body), expected)
	}
}