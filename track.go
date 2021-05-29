package MusicPlayerBackend

type Track struct {
	TrackId int    `json:"-"`
	Title   string `json:"title" binding:"required"`
	Hash    string `json:"-"`
	Genre   string `json:"genre"`
	Artist
	Year     int  `json:"year"`
	HasVideo bool `json:"has_video,string"`
}

type Artist struct {
	ArtistId int    `json:"-"`
	Name     string `json:"artist" binding:"required"`
}
