package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

func JSONMessage(status bool, msg interface{}) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": msg}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := WriteUser(DB, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(JSONMessage(true, id))

	// TODO
	// here will be creating JWT token
}

var JWTAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notAuth := []string{"/auth", "/register"}
		requestPath := r.URL.Path

		// don't check token for this url
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization")

		// token not specified
		if tokenHeader == "" {
			response := JSONMessage(false, "JWT token not specified")
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Add("Content-type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response := JSONMessage(false, "Invalid JWT token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Add("Content-type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}

		tokenPart := splitted[0]
		tk := Token{}

		// parsing token
		token, err := jwt.ParseWithClaims(tokenPart, tk,
			func(token *jwt.Token) (interface{}, error) {
				return []byte("some_example_passwd_string"), nil
			})

		// token is not properly formed
		if err != nil {
			response := JSONMessage(false, "Invalid token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Add("Content-type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}

		if !token.Valid {
			response := JSONMessage(false, "Token is not valid")
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Add("Content-type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}

		// token is valid, performs next handler with context
		log.Printf("UserId: %d", tk.UserId)
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	invalidLogin := func (){
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(JSONMessage(false, "Invalid login or password"))
	}
	user := User{}
	var (
		loginData string
	)

	w.Header().Add("Content-type", "application/json")
	json.NewDecoder(r.Body).Decode(&user)
	if user.Username != "" {
		loginData = user.Username
	} else if user.Email != "" {
		loginData = user.Email
	} else {
		log.Println("Username/Email not specified")
		invalidLogin()
		return
	}

	userDB, ok := GetUserByLogin(loginData)
	if !ok {
		log.Println("user not find in DB")
		invalidLogin()
		return
	}

	if !user.Login(userDB){
		log.Println("wrong login or password")
		invalidLogin()
		return
	}

	json.NewEncoder(w).Encode(userDB)
	return
}
