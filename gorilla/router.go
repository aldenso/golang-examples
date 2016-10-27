package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func notfound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func newrouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/name", getnames).Methods("GET")
	r.HandleFunc("/api/name", postname).Methods("POST")
	r.NotFoundHandler = http.HandlerFunc(notfound)
	return r
}
