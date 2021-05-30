package service

import (
	"github.com/wellWINeo/MusicPlayerBackend"
	"github.com/wellWINeo/MusicPlayerBackend/pkg/repository"
)

type Authorization interface {
	CreateUser(user MusicPlayerBackend.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	UpdateUser(user MusicPlayerBackend.User) error
	DeleteUser(id int) error
	GetUser(id int) (MusicPlayerBackend.User, error)
	SendCode(user MusicPlayerBackend.User) error
	Verify(code int) (MusicPlayerBackend.User, bool)
}

type Playlist interface {

}

type Tracks interface {
	CreateTrack(userId int, track MusicPlayerBackend.Track) (int, error)
	GetTrack(trackId int) (MusicPlayerBackend.Track, error)
	UpdateTrack(track MusicPlayerBackend.Track) error
	DeleteTrack(trackId int) error
	// Upload(userId, trackId int, blob []byte) error
	// Download(userId int) ([]byte, error)
}

type Service struct {
	Authorization
	Playlist
	Tracks
}

func NewService(repos *repository.Repository, mailConfig MailConfig) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization, mailConfig),
		Tracks: NewTracksService(repos.Tracks),
	}
}
