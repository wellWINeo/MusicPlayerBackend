package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type GenreMSSQL struct {
	db *sqlx.DB
}

func NewGenreMSSQL(db * sqlx.DB) *GenreMSSQL {
	return &GenreMSSQL{db: db}
}

func (g *GenreMSSQL) CreateGenre(name string) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s(name) output INSERTED.id_genre values(@p1)",
		genreTable)
	row := g.db.QueryRow(query, name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (g *GenreMSSQL) GetGenreByName(name string) (int, error) {
	var id int
	query := fmt.Sprintf("select id_genre from %s where name=@p1", genreTable)
	err := g.db.Get(&id, query, name)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (g *GenreMSSQL) DeleteGenre(genreId int) error {
	query := fmt.Sprintf("delete from %s where id_genre=@p1", genreTable)
	_, err := g.db.Exec(query, genreId)
	return err
}
