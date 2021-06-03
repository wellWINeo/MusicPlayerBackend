package MusicPlayerBackend

type Track struct {
	TrackId int    `json:"track_id" db:"id_track"`
	Title   string `json:"title" binding:"required" db:"title"`
	Hash    string `json:"-" db:"hash"`
	Genre   string `json:"genre" db:"genre_name"`
	Artist
	Year     int  `json:"year" db:"year"`
	HasVideo bool `json:"has_video" db:"has_video"`
	IsLiked bool `json:"is_liked" db:"is_liked"`
}

type Artist struct {
	ArtistId int    `json:"-"`
	Name     string `json:"artist" binding:"required" db:"name"`
}
