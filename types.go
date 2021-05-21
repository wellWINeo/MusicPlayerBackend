package main

type User struct {
	Id        int
	Username  string
	Passwd    string
	IsPremium bool
	Token     string
}

type Track struct {
	Id     int
	Title  string
	Artist string
}
