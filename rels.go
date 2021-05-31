package MusicPlayerBackend

import "time"


type Like struct {
	LikeId int `json:"-"`
	TrackId int `json:"track_id"`
	UserId int `json:"user_id"`
	Time  time.Time `json:"time"`
}

type History struct {
	HistoryId int `json:"-"`
	TrackId int `json:"track_id" db:"track_id"`
	UserId int `json:"user_id" db:"user_id"`
	Time  time.Time `json:"time" db:"time"`
}

type Referals struct {
	RefId int `json:"-"`
	OldUser int `json:"old_user"`
	NewUser int `json:"new_user"`
}

type Owns struct {
	OwnId int `json:"-"`
	TrackId int `json:"track_id"`
	UserId int `json:"user_id"`
}
