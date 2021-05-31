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
						join %s on id_artist=artist_id
						join %s on id_genre=genre_id
						where id_track=@p1`,
		trackTable, artistsTable, genreTable)
	if err := t.db.Get(&track, query, trackId); err != nil {
		return track, err
	}

	return track, nil
}

func (t *TracksMSSQL) GetAllTracksId(userId int) ([]int, error) {
	response := []int{}

	query := fmt.Sprintf("select id_track from %s where owner_id=@p1", trackTable)
	rows, err := t.db.Query(query, userId)
	if err != nil {
		return response, err
	}

	for rows.Next() {
		var id int

		err := rows.Scan(&id)
		if err != nil {
			return []int{}, err
		}

		response = append(response, id)
	}

	return response, nil
}

func (t *TracksMSSQL) GetAllTracks(userId int) ([]MusicPlayerBackend.Track, error) {
	response := []MusicPlayerBackend.Track{}

	query := fmt.Sprintf(`select id_track, Tracks.title, year, has_video,
						name, Genre.title as genre_name from %s
						join %s on id_artist=artist_id
						join %s on id_genre=genre_id
						where owner_id=@p1`,
		trackTable, artistsTable, genreTable)

	//err := t.db.Get(&response, query, userId)
	err := t.db.Select(&response, query, userId)
	if err != nil {
		return []MusicPlayerBackend.Track{}, err
	}


	return response, nil
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

func (t *TracksMSSQL) SetLike(trackId int) error {
	query := fmt.Sprintf("update %s set is_liked=1-is_liked where id_track=@p1", trackTable)
	_, err := t.db.Exec(query, trackId)
	return err
}

func (t *TracksMSSQL) GetAllLikes(userId int) ([]int, error) {
	var likes []int
	query := fmt.Sprintf("select id_track from %s where owner_id=@p1",
		trackTable)
	err := t.db.Select(&likes, query, userId)
	return likes, err
}
