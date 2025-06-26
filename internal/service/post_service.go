package service

import (
	"context"
	"errors"
	"time"

	"github.com/Fillybodyknow/blog-api/internal/models"
	"github.com/Fillybodyknow/blog-api/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostServiceInterface interface {
	CreatePost(post *models.Post, role string) error
	GetAllPosts() ([]models.Post, error)
	GetAuthorPosts(AuthorID primitive.ObjectID) ([]models.Post, error)
	GetPostsFromTags(tags []string) ([]models.Post, error)
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

func (s *PostService) GetAllPosts() ([]models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var posts []models.Post
	posts, err := s.PostRepository.GetPosts(ctx)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) GetAuthorPosts(authorID primitive.ObjectID) ([]models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	posts, err := s.PostRepository.FindPostByAuthorID(ctx, authorID)
	if err != nil {
		return nil, err
	}
	if len(posts) == 0 {
		return nil, errors.New("ไม่พบโพสต์ของผู้ใช้")
	}
	return posts, nil
}

func (s *PostService) GetPostsFromTags(tags []string) ([]models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	posts, err := s.PostRepository.FindByTags(ctx, tags)
	if err != nil {
		return nil, err
	}
	if len(posts) == 0 {
		return nil, errors.New("ไม่พบโพสต์ตามแท็ก")
	}
	return posts, nil
}
