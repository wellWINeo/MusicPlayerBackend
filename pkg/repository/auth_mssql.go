package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/sirupsen/logrus"
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
	query := fmt.Sprintf("insert into %s(username, email, passwd) output INSERTED.id_user values(:username, :email, :passwd);",
		usersTable)
	result, err := a.db.NamedQuery(query, user)
	result.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (a *AuthMSSQL) GetUser(username, password string) (MusicPlayerBackend.User, error) {
	var user MusicPlayerBackend.User
	query := fmt.Sprintf("select * from %s where username=@p1 and passwd=@p2;", usersTable)
	err := a.db.Get(&user, query, username, password)
	return user, err
}

func (a *AuthMSSQL) GetUserById(id int) (MusicPlayerBackend.User, error) {
	var user MusicPlayerBackend.User
	query := fmt.Sprintf("select * from %s where id_user=@p1", usersTable)
	err := a.db.Get(&user, query, id)
	return user, err
}
