package repository

import (
	"fmt"
	"net/url"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func NewMSSQLDB(cfg Config) (*sqlx.DB, error) {
	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword(cfg.Username, cfg.Password),
		Host:   fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
	}
	db, err := sqlx.Open("sqlserver", u.String())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
