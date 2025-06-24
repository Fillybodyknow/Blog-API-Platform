package service

import (
	"context"
	"errors"

	"github.com/Fillybodyknow/blog-api/internal/models"
	"github.com/Fillybodyknow/blog-api/internal/repository"
	"github.com/Fillybodyknow/blog-api/pkg/utility"
)

type AuthServiceInterface interface {
	Register(ctx context.Context, user *models.User) error
}

type AuthService struct {
	AuthRepository repository.AuthRepositoryInterface
}

func NewAuthService(authRepository repository.AuthRepositoryInterface) *AuthService {
	return &AuthService{AuthRepository: authRepository}
}

func (s *AuthService) Register(ctx context.Context, user *models.User) error {
	exists, _ := s.AuthRepository.FindByEmail(ctx, user.Email)
	if exists != nil {
		return errors.New("email already exists")
	}

	if err := utility.CheckStrongPassword(user.PasswordHash); err != nil {
		return err
	}

	hashing, err := utility.HashPassword(user.PasswordHash)
	if err != nil {
		return err
	}
	user.PasswordHash = hashing
	return s.AuthRepository.InsertUser(ctx, user)
}
