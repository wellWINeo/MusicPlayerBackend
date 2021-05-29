package service

import "github.com/wellWINeo/MusicPlayerBackend/pkg/repository"

type ArtistService struct {
	repo repository.Repository
}

func NewArtistService(repo repository.Repository) *ArtistService {
	return &ArtistService{repo: repo}
}

func (a *ArtistService) CreateArtist(name string) (int, error) {
	artistId, err := a.GetArtistIdByName(name)
	if err == nil {
		return artistId, err
	}

	return a.repo.CreateArtist(name)
}

func (a *ArtistService) GetArtistNameById(artistId int) (string, error) {
	return a.repo.GetArtistById(artistId)
}

func (a *ArtistService) GetArtistIdByName(name string) (int, error) {
	return a.repo.GetArtistByName(name)
}
