package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/wellWINeo/MusicPlayerBackend"
)

type ArtistMSSQL struct {
	db *sqlx.DB
}


func NewArtistMSSQL(db *sqlx.DB) *ArtistMSSQL {
	return &ArtistMSSQL{db: db}
}

// working with Artists
func (t *TracksMSSQL) CreateArtist(artist MusicPlayerBackend.Artist) (int, error) {
	artistDB, err := t.GetArtistByName(artist.Name)
	if err == nil {
		return artistDB.ArtistId, nil
	}

	var id int
	query := fmt.Sprintf("insert into %s(name) ouput INSERTED.id_track values (@p1)", artistsTable)
	row := t.db.QueryRow(query, artist.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (t *TracksMSSQL) GetArtistByName(name string) (MusicPlayerBackend.Artist, error) {
	var artistDB MusicPlayerBackend.Artist
	query := fmt.Sprintf("select id_artist from %s where name=@p1", artistsTable)
	err := t.db.Get(&artistDB, query, name)
	return artistDB, err
}

func (t *TracksMSSQL) GetArtistById(artistId int) (MusicPlayerBackend.Artist, error) {
	var artistDB MusicPlayerBackend.Artist
	query := fmt.Sprintf("select id_artist from %s where artist_id=@p1", artistsTable)
	err := t.db.Get(&artistDB, query, artistId)
	return artistDB, err

}

func (t *TracksMSSQL) DeleteArtist(artistId int) error {
	query := fmt.Sprintf("delete from %s where id_artist=@p1", artistsTable)
	_, err := t.db.Exec(query, artistId)
	return err
}
