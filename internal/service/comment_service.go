package service

import (
	"context"
	"errors"
	"time"

	"github.com/Fillybodyknow/blog-api/internal/models"
	"github.com/Fillybodyknow/blog-api/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentServiceInterface interface {
	Comment(comment string, PostIDStr string, UserID string) error
	EditComment(comment string, PostIDStr string, CommentIDStr string, UserIDStr string) error
	DeleteComment(PostIDStr string, CommentIDStr string, UserIDStr string) error
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

func (s *CommentService) EditComment(comment string, PostIDStr string, CommentIDStr string, UserIDStr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	PostID, _ := primitive.ObjectIDFromHex(PostIDStr)
	UserID, _ := primitive.ObjectIDFromHex(UserIDStr)
	CommentID, _ := primitive.ObjectIDFromHex(CommentIDStr)

	IsOwnerComment, err := s.CommentRepo.GetCommentByID(ctx, PostID, CommentID)
	if err != nil {
		return err
	}

	if IsOwnerComment.UserID != UserID {
		return errors.New("คุณไม่สามารถแก้ไข Comment ของคนอื่นได้ ID คุณคือ " + IsOwnerComment.UserID.Hex() + " ID ของ Comment คือ " + IsOwnerComment.UserID.Hex())
	}

	err = s.CommentRepo.UpdateComment(ctx, comment, PostID, CommentID)
	if err != nil {
		return errors.New("เกิดข้อผิดพลาดในการแก้ไข Comment ขออภัยในความไม่สะดวก")
	}

	return nil
}

func (s *CommentService) DeleteComment(PostIDStr string, CommentIDStr string, UserIDStr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	PostID, _ := primitive.ObjectIDFromHex(PostIDStr)
	UserID, _ := primitive.ObjectIDFromHex(UserIDStr)
	CommentID, _ := primitive.ObjectIDFromHex(CommentIDStr)

	IsOwnerComment, err := s.CommentRepo.GetCommentByID(ctx, PostID, CommentID)
	if err != nil {
		return err
	}

	if IsOwnerComment.UserID != UserID {
		return errors.New("คุณไม่สามารถลบComment ของคนอื่นได้ ID คุณคือ " + IsOwnerComment.UserID.Hex() + " ID ของ Comment คือ " + IsOwnerComment.UserID.Hex())
	}

	err = s.CommentRepo.DeleteComment(ctx, PostID, CommentID)
	if err != nil {
		return errors.New("เกิดข้อผิดพลาดในการลบ Comment ขออภัยในความไม่สะดวก")
	}

	return nil
}
