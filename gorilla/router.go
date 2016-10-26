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
	r.HandleFunc("/api/name/{name}", returnname).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(notfound)
	return r
}
