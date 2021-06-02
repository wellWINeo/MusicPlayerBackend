package repository

import (
	"fmt"
	"net/url"

	"github.com/jmoiron/sqlx"
)

var (
	// tables
	usersTable,
	trackTable,
	genreTable,
	likesTable,
	histroryTable,
	referalsTable,
	ownsTable,
	playlistTable,
	playlistContentTable,
	contentTable,
	trackDataTable,
	artistsTable,
	// stored procedures
	addTrackProc,
	updateTrackProc string
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

func NewMSSQLDB(cfg Config) (*sqlx.DB, error) {
	initTableNames(cfg.DBName)
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

func initTableNames(db_name string) {
	// init tables names
	usersTable = db_name + ".dbo.Users"
	trackTable = db_name + ".dbo.Tracks"
	genreTable = db_name + ".dbo.Genre"
	likesTable = db_name + ".dbo.Likes"
	histroryTable = db_name + ".dbo.History"
	referalsTable = db_name + ".dbo.Referals"
	ownsTable = db_name + ".dbo.Owns"
	playlistTable = db_name + ".dbo.Playlist"
	playlistContentTable = db_name + ".dbo.PlaylistContent"
	contentTable = db_name + ".dbo.PlaylistContent"
	trackDataTable = db_name + ".dbo.TrackData"
	artistsTable = db_name + ".dbo.Artists"
	// init procedures name
	addTrackProc = db_name + ".dbo.AddTrack"
	updateTrackProc = db_name + ".dbo.UpdateTrack"
}
