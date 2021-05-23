package main

import (
	// "log"
	// "strings"

	jwt "github.com/dgrijalva/jwt-go"
)

type User struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Passwd    string `json:"passwd"`
	IsPremium bool   `json:"isPremium"`
}

func (u *User) Login(userDB User) bool {
	return u.Username == userDB.Username || u.Email == userDB.Email &&
		u.Passwd == userDB.Passwd
}

type Track struct {
	Id     int
	Title  string
	Artist string
}

type Token struct {
	UserId uint
	jwt.StandardClaims
}
