package service

import (
	"context"
	"errors"
	"time"

	"github.com/Fillybodyknow/blog-api/internal/models"
	"github.com/Fillybodyknow/blog-api/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LikeServiceInterface interface {
	LikePost(PostIDStr string, userIDStr string) error
	UnlikePost(PostIDStr string, userIDStr string) error
}

type LikeService struct {
	LikeRepository repository.LikeRepositoryInterface
}

func NewLikeService(likeRepository repository.LikeRepositoryInterface) *LikeService {
	return &LikeService{LikeRepository: likeRepository}
}

func (s *LikeService) LikePost(PostIDStr string, userIDStr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	PostID, _ := primitive.ObjectIDFromHex(PostIDStr)
	UserID, _ := primitive.ObjectIDFromHex(userIDStr)

	existingLike, _ := s.LikeRepository.FindLikeByUserID(ctx, PostID, UserID)
	if existingLike != nil {
		return errors.New("คุณได้ Like โพสต์นี้ไปแล้ว")
	}

	like := &models.Like{
		ID:     primitive.NewObjectID(),
		UserID: UserID,
	}
	err := s.LikeRepository.InsertLike(ctx, like, PostID)
	if err != nil {
		return errors.New("เกิดข้อผิดพลาดในการ Like โพสต์ ขออภัยในความไม่สะดวก")
	}
	return nil
}

func (s *LikeService) UnlikePost(PostIDStr string, userIDStr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	PostID, _ := primitive.ObjectIDFromHex(PostIDStr)
	UserID, _ := primitive.ObjectIDFromHex(userIDStr)

	isOwnerLike, _ := s.LikeRepository.FindLikeByUserID(ctx, PostID, UserID)
	if isOwnerLike == nil {
		return errors.New("คุณยังไม่ได้ Like โพสต์นี้")
	}

	err := s.LikeRepository.DeleteLike(ctx, PostID, isOwnerLike.ID)
	if err != nil {
		return errors.New("เกิดข้อผิดพลาดในการ Unlike โพสต์ ขออภัยในความไม่สะดวก -> " + err.Error())
	}
	return nil
}
