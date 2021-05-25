package MusicPlayerBackend

type Track struct {
	TrackId int `json:"-"`
	Title string `json:"title"`
	Hash string `json:"-"`
	Artist
	Year int `json:"year"`
	HasVideo bool `json:"has_video"`
}

type Artist struct {
	ArtistId int `json:"-"`
	Name string `json:"artist"`
}
