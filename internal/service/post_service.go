package service

import (
	"context"
	"errors"
	"time"

	"github.com/Fillybodyknow/blog-api/internal/models"
	"github.com/Fillybodyknow/blog-api/internal/repository"
)

type PostServiceInterface interface {
	CreatePost(post *models.Post, role string) error
}

type PostService struct {
	PostRepository repository.PostRepositoryInterface
	TagRepository  repository.TagRepositoryInterface
}

func NewPostService(postRepository repository.PostRepositoryInterface, tagRepository repository.TagRepositoryInterface) *PostService {
	return &PostService{PostRepository: postRepository, TagRepository: tagRepository}
}

func (s *PostService) CreatePost(post *models.Post, role string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if role != "editor" && role != "admin" {
		return errors.New("คุณไม่สามารถสร้างโพสต์ได้เนื่องจากยังไม่ยืนยันตัวตน")
	}

	if err := s.PostRepository.InsertPost(ctx, post); err != nil {
		return errors.New("ไม่สามารถสร้างโพสต์ได้")
	}

	for _, tagName := range post.Tags {
		existingTag, err := s.TagRepository.FindTagByName(ctx, tagName)
		if err != nil {
			return errors.New("เกิดข้อผิดพลาดในการตรวจสอบแท็ก")
		}
		if existingTag == nil {
			newTag := models.Tag{Name: tagName}
			if err := s.TagRepository.InsertTag(ctx, &newTag); err != nil {
				return errors.New("เกิดข้อผิดพลาดในการสร้างแท็กใหม่")
			}
		}
	}

	return nil
}
