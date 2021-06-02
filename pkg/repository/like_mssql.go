package repository

import (
	"fmt"
	"time"

	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type LikeMSSQL struct {
	db *sqlx.DB
}

func NewLikeMSSQL(db *sqlx.DB) *LikeMSSQL {
	return &LikeMSSQL{db: db}
}

func (l *LikeMSSQL) SetLike(trackId, userId int) error {
	query := fmt.Sprintf("insert into %s values(@p1, @p2, @p3)", likesTable)
	datetime := mssql.DateTime1(time.Now())
	_, err := l.db.Exec(query, trackId, userId, datetime)
	return err
}

func (l *LikeMSSQL) UnsetLike(trackId, userId int) error {
	query := fmt.Sprintf("delete from %s where track_id=@p1 and user_id=@p2",
		likesTable)
	_, err := l.db.Exec(query, trackId, userId)
	return err
}

func (l *LikeMSSQL) GetAll(userId int) ([]int, error) {
	var likes []int
	logrus.Println("here")
	query := fmt.Sprintf("select id_track from %s where owner_id=@p1 and is_liked=1", trackTable)
	if err := l.db.Select(&likes, query, userId); err != nil {
		return []int{}, err
	}

	return likes, nil
}
