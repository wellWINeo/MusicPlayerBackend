package service

import (
	"github.com/wellWINeo/MusicPlayerBackend"
	"github.com/wellWINeo/MusicPlayerBackend/pkg/repository"
)

type TracksService struct {
	repoTracks repository.Tracks
	repoArtist repository.Artist
	repoGenre  repository.Genre
}

func NewTracksService(repoTracks repository.Tracks,
	repoArtist repository.Artist,
	repoGenre repository.Genre) *TracksService {
	return &TracksService{
		repoTracks: repoTracks,
		repoArtist: repoArtist,
		repoGenre:  repoGenre,
	}
}

func (t *TracksService) Create(userId int, track MusicPlayerBackend.Track) (int, error) {
	return 0, nil
}
