package repository

import (
	"fmt"
	"time"

	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
	"github.com/wellWINeo/MusicPlayerBackend"
)

type HistoryMSSQL struct {
	db *sqlx.DB
}

func NewHistoryMSSQL(db *sqlx.DB) *HistoryMSSQL {
	return &HistoryMSSQL{db: db}
}

func (h *HistoryMSSQL) AddHistory(trackId, userId int) error {
	query := fmt.Sprintf("insert into %s values(@p1, @p2, @p3)", histroryTable)
	datetime := mssql.DateTime1(time.Now())
	_, err := h.db.Exec(query, trackId, userId, datetime)
	return err
}

func (h *HistoryMSSQL) GetHistory(userId int) ([]MusicPlayerBackend.History, error) {
	var history []MusicPlayerBackend.History
	query := fmt.Sprintf("select track_id, time from %s where user_id=@p1",
		histroryTable)
	err := h.db.Select(&history, query, userId)
	return history, err
}
