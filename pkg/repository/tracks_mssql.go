package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/wellWINeo/MusicPlayerBackend"
)

type TracksMSSQL struct {
	db *sqlx.DB
}

func NewTracksMSSQL(db *sqlx.DB) *TracksMSSQL {
	return &TracksMSSQL{db: db}
}

func (t *TracksMSSQL) CreateTrack(userId int, track MusicPlayerBackend.Track) (int, error) {
	query := fmt.Sprintf("insert into %s(title, year, has_video)"+
		"output INSERTED.track_id "+
		"values(@p1, @p2, @p3, @p4, @p5)", trackTable)
	row := t.db.QueryRow(query, track.Title,track.Year, track.HasVideo)

	var trackId int
	if err := row.Scan(&trackId); err != nil {
		return 0, err
	}

	query = fmt.Sprintf("insert into %s values(track_id, user_id)", ownsTable)
	_, err := t.db.Exec(query, trackId, userId)
	return trackId, err
}

// TODO
func (t *TracksMSSQL) GetTrack(trackId int) (MusicPlayerBackend.Track, error) {

}

func (t *TracksMSSQL) UpdateTrack(track MusicPlayerBackend.Track) error {

}

func (t *TracksMSSQL) DeleteTrack(trackId int) error {

}
