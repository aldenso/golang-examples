package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
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

func Test_getUserInput(t *testing.T) {
	test := [][]byte{[]byte("N"), []byte("Y"), []byte("s")}
	for _, x := range test {
		tmpfile, err := ioutil.TempFile("", "example")
		if err != nil {
			log.Fatal(err)
		}

		defer os.Remove(tmpfile.Name()) // clean up

		if _, err := tmpfile.Write(x); err != nil {
			log.Fatal(err)
		}

		if _, err := tmpfile.Seek(0, 0); err != nil {
			log.Fatal(err)
		}

		oldStdin := os.Stdin
		defer func() { os.Stdin = oldStdin }() // Restore original Stdin

		os.Stdin = tmpfile
		if err := getUserInput(); err != nil {
			t.Errorf("getUserInput failed: %v", err)
		}

		if err := tmpfile.Close(); err != nil {
			log.Fatal(err)
		}
	}
}
