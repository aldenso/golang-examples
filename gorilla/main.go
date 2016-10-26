package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	router := newrouter()
	log.Fatal(http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, router)))
}
