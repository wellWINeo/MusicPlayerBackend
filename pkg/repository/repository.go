package repository

type Authorization interface {
}

type Playlist interface {
}

type TrackList interface {
}

type Repository struct {
	Authorization
	Playlist
	TrackList
}

func NewRepository() *Repository {
	return &Repository{}
}
