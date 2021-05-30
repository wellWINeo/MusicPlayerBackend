package service

import "github.com/wellWINeo/MusicPlayerBackend/pkg/repository"

type LikeService struct {
	repo repository.Like
}

func NewLikeService(repo repository.Like) *LikeService {
	return &LikeService{repo: repo}
}

func (l *LikeService) SetLike(trackId, userId int) error {
	return l.repo.SetLike(trackId, userId)
}

func (l *LikeService) UnsetLike(trackId, userId int) error {
	return nil
}
