package service

import (
	"github.com/wellWINeo/MusicPlayerBackend"
	"github.com/wellWINeo/MusicPlayerBackend/pkg/repository"
)

type Authorization interface {
	CreateUser(user MusicPlayerBackend.User) (int, error)
	GenerateToken(username, password string) (string, error)
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

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
