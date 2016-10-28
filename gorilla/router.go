package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func notfound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func newrouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/api/user", getusers).Methods("GET")
	r.HandleFunc("/api/user/name/{name}", showuserbyname).Methods("GET")
	r.HandleFunc("/api/user/id/{id}", showuserbyid).Methods("GET")
	r.HandleFunc("/api/user", postuser).Methods("POST")
	r.NotFoundHandler = http.HandlerFunc(notfound)
	return r
}
