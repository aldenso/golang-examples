package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"time"

	"github.com/gorilla/mux"
)

type user struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Createdat time.Time `json:"created_at"`
}

var idCounter int

var memusers []user

func getusers(w http.ResponseWriter, r *http.Request) {
	response, err := json.MarshalIndent(memusers, "", "    ")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func showuserbyname(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userlist := []user{}
	for _, val := range memusers {
		if val.Name == vars["name"] {
			userlist = append(userlist, val)
		}
	}
	if len(userlist) == 0 {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		msg := "user not found"
		response, err := json.MarshalIndent(msg, "", "    ")
		if err != nil {
			panic(err)
		}
		w.Write(response)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		response, err := json.MarshalIndent(userlist, "", "    ")
		if err != nil {
			panic(err)
		}
		w.Write(response)
	}
}

func showuserbyid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userbyid := user{}
	for _, val := range memusers {
		if strconv.Itoa(val.ID) == vars["id"] {
			userbyid = val
		}
	}
	if len(userbyid.Name) == 0 {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		msg := "user not found"
		response, err := json.MarshalIndent(msg, "", "    ")
		if err != nil {
			panic(err)
		}
		w.Write(response)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		response, err := json.MarshalIndent(userbyid, "", "    ")
		if err != nil {
			panic(err)
		}
		w.Write(response)
	}
}

func postuser(w http.ResponseWriter, r *http.Request) {
	var newuser user
	json.NewDecoder(r.Body).Decode(&newuser)
	newuser.Createdat = time.Now()
	newuser.ID = idCounter
	idCounter++
	memusers = append(memusers, newuser)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Location", r.URL.Path+"/id/"+strconv.Itoa(newuser.ID))
	w.WriteHeader(http.StatusCreated)
}
