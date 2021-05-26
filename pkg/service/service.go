package service

import "github.com/wellWINeo/MusicPlayerBackend/pkg/repository"

type Authorization interface {

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
	return &Service{}
}
