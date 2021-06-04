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

type Service struct {
	Authorization
	Playlist
	Tracks
	History
}

func NewService(repos *repository.Repository, mailConfig MailConfig) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization, mailConfig),
		Tracks: NewTracksService(repos.Tracks),
		History: NewHistoryService(repos.History),
		Playlist: NewPlaylistService(repos.Playlist),
	}
}
