package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/wellWINeo/MusicPlayerBackend"
)

type AuthMSSQL struct {
	db *sqlx.DB
}

func NewAuthMSSQL(db *sqlx.DB) *AuthMSSQL {
	return &AuthMSSQL{db}
}

func (a *AuthMSSQL) CreateUser(user MusicPlayerBackend.User) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s(username, email, passwd) values ($1, $2, $3)",
		usersTable)
	row := a.db.QueryRow(query, user.Username, user.Email, user.Password)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (a *AuthMSSQL) GetUser(username, password string) (MusicPlayerBackend.User, error) {
	var user MusicPlayerBackend.User
	query := fmt.Sprintf("select id_user from %s where username=$1 and email=$2", usersTable)
	err := a.db.Get(&user, query, username, password)
	return user, err
}
