package service

import (
	"context"
	"time"

	"github.com/Fillybodyknow/blog-api/internal/models"
	"github.com/Fillybodyknow/blog-api/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentServiceInterface interface {
	Comment(comment string, PostIDStr string, UserID string) error
}

type CommentService struct {
	CommentRepo repository.CommentRepositoryInterface
}

func NewCommentService(commentRepo repository.CommentRepositoryInterface) *CommentService {
	return &CommentService{CommentRepo: commentRepo}
}

func (s *CommentService) Comment(comment string, PostIDStr string, UserIDStr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	PostID, _ := primitive.ObjectIDFromHex(PostIDStr)
	UserID, _ := primitive.ObjectIDFromHex(UserIDStr)

	commentModel := models.Comment{
		ID:        primitive.NewObjectID(),
		Content:   comment,
		UserID:    UserID,
		CreatedAt: time.Now(),
	}
	err := s.CommentRepo.InsertComment(ctx, &commentModel, PostID)
	return err
}
