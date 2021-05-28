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

type TrackList interface {

}

type Service struct {
	Authorization
	Playlist
	TrackList
}

func NewService(repos *repository.Repository, mailConfig MailConfig) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization, mailConfig),
	}
}
