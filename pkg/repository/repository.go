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

type Tracks interface {
	CreateTrack(userId int, track MusicPlayerBackend.Track) (int, error)
	GetTrack(trackId int) (MusicPlayerBackend.Track, error)
	GetAllTracksId(userId int) ([]int, error)
	GetAllTracks(userId int) ([]MusicPlayerBackend.Track, error)
	UpdateTrack(track MusicPlayerBackend.Track) error
	DeleteTrack(trackId int) error
	// UploadTrack()
	// DownloadTrack()
}

type Like interface {
	SetLike(trackId, userId int) error
	GetAll(userId int) ([]int, error)
	UnsetLike(trackId, userId int) error
}

type Repository struct {
	Authorization
	Playlist
	Like
	Tracks
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMSSQL(db),
		Tracks: NewTracksMSSQL(db),
		Like: NewLikeMSSQL(db),
	}
}
