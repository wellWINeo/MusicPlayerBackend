package main

import (
	//"fmt"
	"net/http"
	"log"
)

func main() {
	// configuring routing
	mux := http.NewServeMux()

	// register handler
	mux.HandleFunc("/register", func (w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-type", "text/plain")
		w.Write([]byte("Register handler"))
	})


	// auth handler
	mux.HandleFunc("/auth", func (w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-type", "text/plain")
		w.Write([]byte("Auth handler"))
	})

	// user handler
	mux.HandleFunc("/user", func (w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-type", "text/plain")
		if r.Method != http.MethodGet {
			w.Write([]byte("Go fuck yourself"))
		} else {
			w.Write([]byte("Auth handler"))
		}
	})

	// server started
	log.Println("Server started...")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
