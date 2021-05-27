package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/wellWINeo/MusicPlayerBackend"
)

type Authorization interface {
	CreateUser(user MusicPlayerBackend.User) (int, error)
	GetUser(username, password string) (MusicPlayerBackend.User, error)
}

type Playlist interface {
}

type TrackList interface {
}

type Repository struct {
	Authorization
	Playlist
	TrackList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMSSQL(db),
	}
}
