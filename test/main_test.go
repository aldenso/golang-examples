package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_getURL(t *testing.T) {
	ts1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if r.Method != "GET" {
			t.Errorf("Expected 'GET' request, got '%s'", r.Method)
		}
		if r.URL.RequestURI() != "/myfakeuri" {
			t.Errorf("Expected request to '/myfakeuri', got '%s'", r.URL.RequestURI())
		}
	}))
	defer ts1.Close()
	_, err := getURL(ts1.URL + "/myfakeuri")
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	// Test for expected error
	_, err = getURL("http://localhost/myfakeuri")
	if err == nil {
		t.Errorf("Expected failure, got unexpected success")
	}
}
