package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/wellWINeo/MusicPlayerBackend"
	// "github.com/wellWINeo/MusicPlayerBackend"
)

type PlaylistMSSQL struct {
	db *sqlx.DB
}

func NewPlaylistMSSQL(db *sqlx.DB) *PlaylistMSSQL {
	return &PlaylistMSSQL{db: db}
}


func (p *PlaylistMSSQL) CreatePlaylist(title string, userId int) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s output INSERTED.id_playlist values (@p1, @p2)",
		playlistTable)
	logrus.Println(title)
	row := p.db.QueryRow(query, userId, title)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (p *PlaylistMSSQL) GetPlaylist(id int) ([]MusicPlayerBackend.Track, error) {
	var tracks []MusicPlayerBackend.Track
	query := fmt.Sprintf(`select id_track, Tracks.title, year, has_video,
						  name, Genre.title as genre_name from %s
						  join %s on track_id=id_track
						  join %s on artist_id=id_artist
 						  join %s on genre_id=id_genre
						  where playlist_id=@p1`,
		playlistContentTable, trackTable, artistsTable, genreTable)
	err := p.db.Select(&tracks, query, id)
	return tracks, err
}

func (p *PlaylistMSSQL) UpdatePlaylist(title string, playlistId int) error {
	query := fmt.Sprintf("update %s set title=@p1 where id_playlist=@p2", playlistTable)
	_, err := p.db.Exec(query, title, playlistId)
	return err
}

func (p *PlaylistMSSQL) DeletePlaylist(id int) error {
	query := fmt.Sprintf("delete from %s where id_playlist=@p1", playlistTable)
	_, err := p.db.Exec(query, id)
	return err
}

func (p *PlaylistMSSQL) AddToPlaylist(playlistId, trackId int) error {
	query := fmt.Sprintf("insert into %s values(@p1, @p2)", playlistContentTable)
	_, err := p.db.Exec(query, trackId, playlistId)
	return err
}

func (p *PlaylistMSSQL) RemoveFromPlaylist(playlistId, trackId int) error {
	query := fmt.Sprintf("delete from %s where playlist_id=@p1 and track_id=@p2",
		playlistContentTable)
	_, err := p.db.Exec(query, trackId, playlistId)
	return err
}

func (p *PlaylistMSSQL) GetUsersPlaylists(userId int) ([]MusicPlayerBackend.Playlist, error) {
	var playlists []MusicPlayerBackend.Playlist
	query := fmt.Sprintf("select id_playlist, title from %s where user_id=@p1",
		playlistTable)
	err := p.db.Select(&playlists, query, userId)
	return playlists, err
}
