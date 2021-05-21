package main

type User struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Passwd    string `json:"passwd"`
	IsPremium bool   `json:"isPremium"`
}

type Track struct {
	Id     int
	Title  string
	Artist string
}
