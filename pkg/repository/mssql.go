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

func initTableNames(dbName string) {
	// init tables names
	usersTable = dbName + ".dbo.Users"
	trackTable = dbName + ".dbo.Tracks"
	genreTable = dbName + ".dbo.Genre"
	likesTable = dbName + ".dbo.Likes"
	histroryTable = dbName + ".dbo.History"
	referalsTable = dbName + ".dbo.Referals"
	ownsTable = dbName + ".dbo.Owns"
	playlistTable = dbName + ".dbo.Playlist"
	playlistContentTable = dbName + ".dbo.PlaylistContent"
	contentTable = dbName + ".dbo.PlaylistContent"
	trackDataTable = dbName + ".dbo.TrackData"
	artistsTable = dbName + ".dbo.Artists"
	// init procedures name
	addTrackProc = dbName + ".dbo.AddTrack"
	updateTrackProc = dbName + ".dbo.UpdateTrack"
}
