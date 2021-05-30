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
	query := fmt.Sprintf("exec %s @p1, @p2, @p3, @p4, @p5, @p6", addTrackProc)
	row := t.db.QueryRow(query, track.Title, track.Artist.Name, track.Genre,
		track.Year, track.HasVideo, userId)

	var trackId int
	if err := row.Scan(&trackId); err != nil {
		return 0, err
	}

	return trackId, nil
}

func (t *TracksMSSQL) GetTrack(trackId int) (MusicPlayerBackend.Track, error) {
	var track MusicPlayerBackend.Track
	query := fmt.Sprintf(`select id_track, Tracks.title, year, has_video,
						name, Genre.title as genre_name from %s
						join Artists on id_artist=artist_id
						join Genre on id_genre=genre_id
						where id_track=@p1`,
		trackTable)
	row := t.db.QueryRow(query, trackId)
	if err := row.Scan(&track); err != nil {
		return track, err
	}

	return track, nil
}

func (t *TracksMSSQL) UpdateTrack(track MusicPlayerBackend.Track) error {
	query := fmt.Sprintf("exec %s @p1, @p2, @p3, @p4, @p5, @p6", updateTrackProc)
	_, err := t.db.Exec(query, track.TrackId, track.Title, track.Name, track.Genre,
		track.Year, track.HasVideo)
	return err
}

func (t *TracksMSSQL) DeleteTrack(trackId int) error {
	query := fmt.Sprintf("delete from %s where id_track=@p1", trackTable)
	_, err := t.db.Exec(query, trackId)
	return err
}
