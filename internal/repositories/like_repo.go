package repositories

import (
	"blog/internal/models"

	"gorm.io/gorm"
)

type LikeRepository struct {
	DB *gorm.DB
}

func NewLikeRepository(db *gorm.DB) *LikeRepository {
	return &LikeRepository{DB: db}
}

func (r *LikeRepository) AddLike(like *models.Like) error {
	return r.DB.Create(like).Error
}

func (r *LikeRepository) RemoveLike(postID, userID uint) error {
	return r.DB.Where("post_id = ? AND user_id = ?", postID, userID).Delete(&models.Like{}).Error
}

func (r *LikeRepository) GetLikesCount(postID uint) (int64, error) {
	var count int64
	err := r.DB.Model(&models.Like{}).Where("post_id = ?", postID).Count(&count).Error
	return count, err
}
