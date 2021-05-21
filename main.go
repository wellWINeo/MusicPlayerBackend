package main

import (
	//"fmt"
	"database/sql"
	"log"
	"net/http"

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
	mux := http.NewServeMux()

	// register handler
	// mux.HandleFunc("/register", func (w http.ResponseWriter, r *http.Request){
	// 	w.Header().Set("Content-type", "text/plain")
	// 	w.Write([]byte("Register handler"))
	// })
	mux.HandleFunc("/register", RegisterHandler)

	// auth handler
	mux.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "text/plain")
		w.Write([]byte("Auth handler"))
	})

	// user handler
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
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
