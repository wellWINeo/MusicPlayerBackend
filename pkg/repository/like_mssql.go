package repository

import (
	"fmt"
	"time"

	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
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
	return nil
}
