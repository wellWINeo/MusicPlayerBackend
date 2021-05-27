package repository

import (
	"fmt"
	"net/url"

	"github.com/jmoiron/sqlx"
)

var (
	usersTable = "Users"
	tracksTable = "Tracks"
	genreTable = "Genre"
	likesTable = "Likes"
	historyTable = "History"
	referalsTable = "Referals"
	ownsTable = "Owns"
	playlistTable = "Playlist"
	contentTable = "PlaylistContent"
	artistsTable = "Artists"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

func NewMSSQLDB(cfg Config) (*sqlx.DB, error) {
	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword(cfg.Username, cfg.Password),
		Host:   fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
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
