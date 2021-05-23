package main

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// funtction to create new user and write it to database
func WriteUser(db *sql.DB, user User) (uint, error) {
	result, err := db.Exec("insert into users(login, passwd, is_premium) values($1, $2, $3)",
		user.Username, user.Passwd, user.IsPremium)
	if err != nil {
		log.Println(err)
		return *new(uint), err
	}
	res, err := result.LastInsertId()
	log.Println(res)
	log.Println(err)
	return uint(res), err
}

// function to get user from database
func GetUserById(db *sql.DB, id string) (user User, err error) {
	rows, err := db.Query("select * from users where user_id=$1", id)
	if err != nil {
		return user, err
	}

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Passwd, &user.IsPremium)
		if err != nil {
			return user, err
		}
	}
	return user, nil
}

func GetUserByLogin(login string) (user User, ok bool) {
	var (
		rows *sql.Rows
		err error
	)
	if strings.Contains(login, "@") {
		rows, err = DB.Query("select * from users where email=$1", login)
	} else {
		rows, err = DB.Query("select * from users where login=$1", login)
	}
	defer rows.Close()
	if err != nil {
		log.Println(err)
		return user, false
	}

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Passwd, &user.IsPremium)
		if err != nil {
			log.Println(err)
			return user, false
		}
	}
	return user, true
}

// function to update user info
func UpdateUser(db *sql.DB, token string, user User) error {
	oldUser, err := GetUserById(db, token)
	if err != nil {
		return err
	}

	// login changed
	if user.Username != oldUser.Username {
		_, err := db.Exec("update users set login=$1 where token=$2", user.Username, token)
		if err != nil {
			return err
		}
	}

	// password changed
	if user.Passwd != oldUser.Passwd {
		_, err := db.Exec("update users set passwd=$1 where token=$2", user.Passwd, token)
		if err != nil {
			return err
		}
	}

	// premium status changed
	if user.IsPremium != oldUser.IsPremium {
		_, err := db.Exec("update users set is_premium=$1 where token=$2", user.IsPremium, token)
		if err != nil {
			return err
		}
	}

	return nil
}
