package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func returnname(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response, err := json.MarshalIndent(vars, "", "    ")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
