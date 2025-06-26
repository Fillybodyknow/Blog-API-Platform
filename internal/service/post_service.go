package service

import (
	"context"
	"errors"
	"strings"
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
	GetPostByID(id primitive.ObjectID) (*models.Post, error)
	EditMePost(post *models.Post, role string, UserIDStr string, PostIDStr string) error
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

	if err := s.PostRepository.Insert(ctx, post); err != nil {
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
	posts, err := s.PostRepository.Get(ctx)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) GetAuthorPosts(authorID primitive.ObjectID) ([]models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	posts, err := s.PostRepository.FindByAuthorID(ctx, authorID)
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

func (s *PostService) GetPostByID(id primitive.ObjectID) (*models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	post, err := s.PostRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, errors.New("ไม่พบโพสต์")
	}
	return post, nil
}

func (s *PostService) EditMePost(post *models.Post, role string, UserIDStr string, PostIDStr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	UserID, err := primitive.ObjectIDFromHex(UserIDStr)
	if err != nil {
		return err
	}

	PostID, err := primitive.ObjectIDFromHex(PostIDStr)
	if err != nil {
		return err
	}

	PostByID, err := s.PostRepository.FindByID(ctx, PostID)
	if err != nil {
		return err
	}
	if PostByID == nil {
		return errors.New("ไม่พบโพสต์ที่ต้องการแก้ไข")
	}

	if PostByID.AuthorID != UserID {
		return errors.New("คุณไม่สามารถแก้ไขโพสต์ของผู้อื่นได้")
	}

	if role != "editor" && role != "admin" {
		return errors.New("คุณไม่สามารถแก้ไขโพสต์ได้เนื่องจากยังไม่ยืนยันตัวตน")
	}

	if err := s.PostRepository.Update(ctx, PostID, post); err != nil {
		return errors.New("ไม่สามารถแก้ไขโพสต์ได้ ขออภัยในความไม่สะดวก")
	}

	for _, tagName := range post.Tags {
		tagName = strings.TrimSpace(tagName)

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
