package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {
	// Create a request to the root endpoint
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler directly
	RootHandler(rr, req)

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

func TestHiHandler(t *testing.T) {
	// Create a request to the /hi endpoint
	req, err := http.NewRequest("GET", "/hi", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler directly
	HiHandler(rr, req)

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