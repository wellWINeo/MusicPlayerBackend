package service

import (
	"github.com/wellWINeo/MusicPlayerBackend"
	"github.com/wellWINeo/MusicPlayerBackend/pkg/repository"
)

type TracksService struct {
	repo repository.Tracks
}

func NewTracksService(repo repository.Tracks) *TracksService {
	return &TracksService{repo: repo}
}

func (t *TracksService) CreateTrack(userId int, track MusicPlayerBackend.Track) (int, error) {
	return t.repo.CreateTrack(userId, track)
}

func (t *TracksService) GetTrack(trackId int) (MusicPlayerBackend.Track, error) {
	return t.repo.GetTrack(trackId)
}

func (t *TracksService) UpdateTrack(track MusicPlayerBackend.Track) error {
	return t.repo.UpdateTrack(track)
}

func (t *TracksService) DeleteTrack(trackId int) error {
	return t.repo.DeleteTrack(trackId)
}

func (t *TracksService) GetAllTracks(userId int) ([]MusicPlayerBackend.Track, error){
	return t.repo.GetAllTracks(userId)
}

func (t *TracksService) GetAllTracksId(userId int) ([]int, error){
	return t.repo.GetAllTracksId(userId)
}

func (t *TracksService) SetLike(trackId int) error {
	return t.repo.SetLike(trackId)
}

func (t *TracksService) GetAllLikes(userId int) ([]int, error) {
	return t.repo.GetAllLikes(userId)
}
