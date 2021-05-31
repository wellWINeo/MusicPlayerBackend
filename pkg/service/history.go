package service

import (
	"github.com/wellWINeo/MusicPlayerBackend"
	"github.com/wellWINeo/MusicPlayerBackend/pkg/repository"
)

type HistoryService struct {
	repo repository.History
}

func NewHistoryService(repo repository.History) *HistoryService {
	return &HistoryService{repo: repo}
}

func (h *HistoryService) AddHistory(trackId, userId int) error {
	return h.repo.AddHistory(trackId, userId)
}

func (h *HistoryService) GetHistory(userId int) ([]MusicPlayerBackend.History, error) {
	return h.repo.GetHistory(userId)
}
