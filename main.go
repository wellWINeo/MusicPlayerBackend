package main

import (
	//"fmt"
	"database/sql"
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var (
	DB *sql.DB
	connString string
)

func main() {
	var err error
	connString = "music.db"
	DB, err = sql.Open("sqlite3", connString)
	if err != nil {
		log.Fatal(err)
	}
	// configuring routing
	router := mux.NewRouter()

	// middleware
	router.Use(JWTAuthentication)

	// register handler
	router.HandleFunc("/register", RegisterHandler).Methods("POST")

	// auth handler
	router.HandleFunc("/auth", AuthHandler).Methods("POST")

	// user handler - GET
	router.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "text/plain")
		if r.Method != http.MethodGet {
			w.Write([]byte("Go fuck yourself"))
		} else {
			w.Write([]byte("Auth handler"))
		}
	}).Methods("GET")

	// user handler - POST
	router.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-type", "text/plain")
		if r.Method != http.MethodGet {
			w.Write([]byte("Go fuck yourself"))
		} else {
			w.Write([]byte("Auth handler"))
		}
	}).Methods("POST")

	// server started
	log.Println("Server started...")
	log.Fatal(http.ListenAndServe(":8000", router))

}
