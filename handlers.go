package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println(user)
	id, err := WriteUser(DB, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	log.Println(id)
	w.Write([]byte(string(id)))
	// TODO
	// here will be creating JWT token
}
