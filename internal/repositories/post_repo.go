package repositories

import (
	"blog/internal/models"

	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{DB: db}
}

func (r *PostRepository) CreatePost(post *models.Post) error {
	return r.DB.Create(post).Error
}

func (r *PostRepository) GetPostsByUserID(userID uint) ([]models.Post, error) {
	var posts []models.Post
	err := r.DB.Where("user_id = ?", userID).Find(&posts).Error
	return posts, err
}

func (r *PostRepository) DeletePost(postID uint) error {
	return r.DB.Where("id = ?", postID).Delete(&models.Post{}).Error
}
