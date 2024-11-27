package services

import (
	"blog/internal/models"
	"blog/internal/repositories"
)

type PostService struct {
	PostRepo *repositories.PostRepository
}

func NewPostService(postRepo *repositories.PostRepository) *PostService {
	return &PostService{PostRepo: postRepo}
}

func (s *PostService) CreatePost(post *models.Post) error {
	return s.PostRepo.CreatePost(post)
}

func (s *PostService) GetPostsByUserID(userID uint) ([]models.Post, error) {
	return s.PostRepo.GetPostsByUserID(userID)
}

func (s *PostService) DeletePost(postID uint) error {
	return s.PostRepo.DeletePost(postID)
}
