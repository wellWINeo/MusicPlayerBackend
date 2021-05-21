package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)


type Store struct {
	Conn sql.DB
}

// returns new store instance
func NewStore(connString string) *Store {
	response, err := sql.Open("sqlite3", connString)
	if err != nil {
		log.Fatal(err)
	}
	return &Store{Conn: *response}
}

// method to create new user and write it to database
func (s *Store) WriteUser(user User) error {
	_, err := s.Conn.Exec("insert into users(login, passwd, is_premium) values($1, $2, $3)",
		user.Username, user.Passwd, user.IsPremium)
	return err
}

// method to get user from database
func (s *Store) GetUser(token string) (user User, err error){
	rows, err := s.Conn.Query("select * from users where token=$1", token)
	if err != nil {
		return user, err
	}

	for rows.Next(){
		err := rows.Scan(&user.Id, &user.Username, &user.Passwd, &user.Token, &user.IsPremium)
		if err != nil {
			return user, err
		}
	}
	return user, nil
}

// method to update user info
func (s *Store) UpdateUser(token string, user User) error {
	oldUser, err := s.GetUser(token)
	if err != nil {
		return err
	}

	// login changed
	if (user.Username != oldUser.Username){
		_, err := s.Conn.Exec("update users set login=$1 where token=$2", user.Username, token)
		if err != nil {
			return err
		}
	}

	// password changed
	if (user.Passwd != oldUser.Passwd){
		_, err := s.Conn.Exec("update users set passwd=$1 where token=$2", user.Passwd, token)
		if err != nil {
			return err
		}
	}

	// premium status changed
	if (user.IsPremium != oldUser.IsPremium){
		_, err := s.Conn.Exec("update users set is_premium=$1 where token=$2", user.IsPremium, token)
		if err != nil {
			return err
		}
	}

	// token changed
	if (user.Token != oldUser.Token){
		_, err := s.Conn.Exec("update users set token=$1 where token=$2", user.Token, token)
		if err != nil {
			return err
		}
	}




}
