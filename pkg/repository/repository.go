package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
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
	return &Repository{}
}
