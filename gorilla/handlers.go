package main

import (
	"encoding/json"
	"net/http"

	"time"
)

type name struct {
	Name      string    `json:"name"`
	Createdat time.Time `json:"created_at"`
}

var memnames []name

func getnames(w http.ResponseWriter, r *http.Request) {
	response, err := json.MarshalIndent(memnames, "", "    ")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func postname(w http.ResponseWriter, r *http.Request) {
	var newname name
	json.NewDecoder(r.Body).Decode(&newname)
	newname.Createdat = time.Now()
	memnames = append(memnames, newname)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Location", r.URL.Path+"/"+newname.Name)
	w.WriteHeader(http.StatusCreated)
}
