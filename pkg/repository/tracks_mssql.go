package repository

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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

func (t *TracksMSSQL) UpdateTrack(trackId int, track MusicPlayerBackend.Track) error {
	query := fmt.Sprintf("exec %s @p1, @p2, @p3, @p4, @p5, @p6", updateTrackProc)
	_, err := t.db.Exec(query, trackId, track.Title, track.Name, track.Genre,
		track.Year, track.HasVideo)
	logrus.Printf("exec %s %d, %s, %s, %s, %d, %t", updateTrackProc, trackId,
		track.Title, track.Name, track.Genre, track.Year, track.HasVideo)
	return err
}

func (t *TracksMSSQL) DeleteTrack(trackId int) error {
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("delete from %s where track_id=@p1", histroryTable)
	_, err = tx.Exec(query, trackId)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = fmt.Sprintf("delete from %s where id_track=@p1", trackTable)
	_, err = tx.Exec(query, trackId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (t *TracksMSSQL) SetLike(trackId int) error {
	query := fmt.Sprintf("update %s set is_liked=1-is_liked where id_track=@p1", trackTable)
	_, err := t.db.Exec(query, trackId)
	return err
}

func (t *TracksMSSQL) GetAllLikes(userId int) ([]int, error) {
	var likes []int
	query := fmt.Sprintf("select id_track from %s where owner_id=@p1 and is_liked=1",
		trackTable)
	err := t.db.Select(&likes, query, userId)
	return likes, err
}

func (t *TracksMSSQL) UploadTrack(trackId int, blob []byte) error {
	var id int
	hashBytes := sha1.Sum(blob)
	hash := hex.EncodeToString(hashBytes[:])
	query := fmt.Sprintf("insert into %s output INSERTED.id_track_data values(@p1, @p2)",
		trackDataTable)
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}

	row := tx.QueryRow(query, hash, blob)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return err
	}

	query = fmt.Sprintf("update %s set data=@p1 where id_track=@p2", trackTable)
	_, err = tx.Exec(query, id, trackId)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (t *TracksMSSQL) DownloadTrack(trackId int) ([]byte, error) {
	var blob []byte

	query := `select convert (varbinary, "MusicPlayer"."dbo"."TrackData"."data")
 			   from "MusicPlayer"."dbo"."Tracks" join
			  "MusicPlayer"."dbo"."TrackData" on
			 "MusicPlayer"."dbo"."Tracks".data=id_track_data where id_track=1`
	row := t.db.QueryRow(query, trackId)
	if err := row.Scan(&blob); err != nil {
		return blob, err
	}

	return blob, nil
}
