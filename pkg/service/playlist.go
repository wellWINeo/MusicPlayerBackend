package service

import (
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/wellWINeo/MusicPlayerBackend"
	"github.com/wellWINeo/MusicPlayerBackend/pkg/repository"
)

type PlaylistService struct {
	repo repository.Playlist
}

func NewPlaylistService(repo repository.Playlist) *PlaylistService {
	return &PlaylistService{repo: repo}
}

func (p *PlaylistService) CreatePlaylist(title string, userId int) (int, error) {
	alreadyExists, err := p.GetUsersPlaylists(userId)
	if err != nil {
		return 0, err
	}

	for _, value := range alreadyExists {
		if value.Title == title {
			logrus.Printf("%s = %s, %t", value.Title, title, value.Title == title)
			return 0, errors.New("playlist with this title already exist")
		}
	}

	// insert to DB
	return p.repo.CreatePlaylist(title, userId)
}

func (p *PlaylistService) UpdatePlaylist(title string, playlistId int) error {
	return p.repo.UpdatePlaylist(title, playlistId)
}

func (p *PlaylistService) DeletePlaylist(id int) error {
	return p.repo.DeletePlaylist(id)
}

func (p *PlaylistService) AddToPlaylist(playlistId, trackId int) error {
	return p.repo.AddToPlaylist(playlistId, trackId)
}

func (p *PlaylistService) RemoveFromPlaylist(playlistId, trackId int) error {
	return p.repo.RemoveFromPlaylist(playlistId, trackId)
}

func (p *PlaylistService) GetPlaylist(id int) ([]MusicPlayerBackend.Track, error) {
	return p.repo.GetPlaylist(id)
}

func (p *PlaylistService) GetUsersPlaylists(userId int) ([]MusicPlayerBackend.Playlist, error) {
	return p.repo.GetUsersPlaylists(userId)
}
