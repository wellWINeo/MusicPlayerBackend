package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/wellWINeo/MusicPlayerBackend"
)

type Authorization interface {
	CreateUser(user MusicPlayerBackend.User) (int, error)
	GetUser(username, password string) (MusicPlayerBackend.User, error)
	GetUserById(id int) (MusicPlayerBackend.User, error)
	DeleteUser(id int) error
	UpdateUser(user MusicPlayerBackend.User) error
}

type Playlist interface {

}

type Artist interface {
	CreateArtist(name string) (int, error)
	GetArtistNameById(artistId int) (string, error)
	GetArtistIdByName(name string) (int, error)
}

type Genre interface {
	CreateGenre(name string) (int, error)
	GetGenreByName(name string) (int, error)
	DeleteGenre(genreId int) error
}

type Tracks interface {
	CreateTrack(userId int, track MusicPlayerBackend.Track) (int, error)
	// Get()
	// Update()
	// Delete()
	// UploadTrack()
	// DownloadTrack()
}

type Repository struct {
	Authorization
	Playlist
	Tracks
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMSSQL(db),
	}
}
