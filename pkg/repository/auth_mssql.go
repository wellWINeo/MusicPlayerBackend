package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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
	query := fmt.Sprintf("insert into %s(username, email, passwd) output INSERTED.id_user values(@p1, @p2, @p3);",
		usersTable)
	row := a.db.QueryRow(query, user.Username, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (a *AuthMSSQL) CreateReferal(oldUser, newUser int) error {
	query := fmt.Sprintf("insert into %s values(@p1, @p2)", referalsTable)
	logrus.Println(query)
	_, err := a.db.Exec(query, oldUser, newUser)
	return err
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

func (a *AuthMSSQL) DeleteUser(id int) error {
	// deleting referals records
	query := fmt.Sprintf("update %s set old_user_id=NULL where old_user_id=@p1",
		referalsTable)
	_, err := a.db.Exec(query, id)
	if err != nil {
		return err
	}
	query = fmt.Sprintf("update %s set new_user_id=NULL where new_user_id=@p1",
		referalsTable)
	_, err = a.db.Exec(query, id)

	// deleting user
	query = fmt.Sprintf("delete from %s where id_user=@p1", usersTable)
	_, err = a.db.Exec(query, id)
	return err
}

func (a *AuthMSSQL) UpdateUser(user MusicPlayerBackend.User) error {
	query := fmt.Sprintf("update %s set username=@p1, email=@p2, "+
		"is_premium=@p4, is_verified=@p5", usersTable)
	_, err := a.db.Exec(query, user.Username, user.Email, user.Password,
		user.IsPremium, user.IsVerified)
	return err
}

func (a *AuthMSSQL) BuyPremium(userId int) error {
	query := fmt.Sprintf("update %s set is_premium=1 where id_user=@p1",
		usersTable)
	_, err := a.db.Exec(query, userId)
	return err
}
