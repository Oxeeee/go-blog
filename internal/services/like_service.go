package services

import (
	"blog/internal/models"
	"blog/internal/repositories"
)

type LikeService struct {
	LikeRepo *repositories.LikeRepository
}

func NewLikeService(likeRepo *repositories.LikeRepository) *LikeService {
	return &LikeService{LikeRepo: likeRepo}
}

func (s *LikeService) AddLike(postID, userID uint) error {
	like := &models.Like{
		PostID: postID,
		UserID: userID,
	}
	return s.LikeRepo.AddLike(like)
}

func (s *LikeService) RemoveLike(postID, userID uint) error {
	return s.LikeRepo.RemoveLike(postID, userID)
}

func (s *LikeService) GetLikesCount(postID uint) (int64, error) {
	return s.LikeRepo.GetLikesCount(postID)
}
