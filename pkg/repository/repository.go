package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/wellWINeo/MusicPlayerBackend"
)

type Authorization interface {
	CreateUser(user MusicPlayerBackend.User) (int, error)
	CreateReferal(oldUser, newUser int) error
	GetUser(username, password string) (MusicPlayerBackend.User, error)
	GetUserById(id int) (MusicPlayerBackend.User, error)
	DeleteUser(id int) error
	UpdateUser(user MusicPlayerBackend.User) error
	BuyPremium(userId int) error
}

type Playlist interface {
	CreatePlaylist(title string, userId int) (int, error)
	GetPlaylist(id int) ([]MusicPlayerBackend.Track, error)
	UpdatePlaylist(title string, playlistId int) error
	DeletePlaylist(id int) error
	AddToPlaylist(playlistId, trackId int) error
	RemoveFromPlaylist(playlistId, trackId int) error
	GetUsersPlaylists(userId int) ([]MusicPlayerBackend.Playlist, error)
}

type Tracks interface {
	CreateTrack(userId int, track MusicPlayerBackend.Track) (int, error)
	GetTrack(trackId int) (MusicPlayerBackend.Track, error)
	GetAllTracksId(userId int) ([]int, error)
	GetAllTracks(userId int) ([]MusicPlayerBackend.Track, error)
	UpdateTrack(trackId int, track MusicPlayerBackend.Track) error
	DeleteTrack(trackId int) error
	UploadTrack(trackId int, blob []byte) error
	DownloadTrack(trackId int) ([]byte, error)
	SetLike(trackId int) error
	GetAllLikes(userId int) ([]int, error)
}

type History interface {
	AddHistory(trackId, userId int) error
	GetHistory(userId int) ([]MusicPlayerBackend.History, error)
}

type Repository struct {
	Authorization
	Playlist
	History
	Tracks
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthMSSQL(db),
		Tracks:        NewTracksMSSQL(db),
		History:       NewHistoryMSSQL(db),
		Playlist:      NewPlaylistMSSQL(db),
	}
}
