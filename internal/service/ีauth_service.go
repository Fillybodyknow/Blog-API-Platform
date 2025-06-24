package service

import (
	"context"
	"errors"

	"github.com/Fillybodyknow/blog-api/internal/models"
	"github.com/Fillybodyknow/blog-api/internal/repository"
	"github.com/Fillybodyknow/blog-api/pkg/utility"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	Register(ctx context.Context, user *models.User) error
	Login(ctx context.Context, username string, password string) (models.User, string, error)
}

type AuthService struct {
	AuthRepository repository.AuthRepositoryInterface
}

func NewAuthService(authRepository repository.AuthRepositoryInterface) *AuthService {
	return &AuthService{AuthRepository: authRepository}
}

func (s *AuthService) Register(ctx context.Context, user *models.User) error {
	Emailexists, _ := s.AuthRepository.FindByEmailOrUsername(ctx, user.Email)
	if Emailexists != nil {
		return errors.New("email นี้ถูกใช้แล้ว")
	}
	Usernameexists, _ := s.AuthRepository.FindByEmailOrUsername(ctx, user.Username)
	if Usernameexists != nil {
		return errors.New("username นี้ถูกใช้แล้ว")
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

func (s *AuthService) Login(ctx context.Context, username string, password string) (models.User, string, error) {
	user, _ := s.AuthRepository.FindByEmailOrUsername(ctx, username)
	if user == nil {
		return models.User{}, "", errors.New("ไม่พบบัญชีผู้ใช้")
	}

	// เช็กรหัสผ่าน
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return models.User{}, "", errors.New("รหัสผ่านไม่ถูกต้อง")
	}

	token, err := utility.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return models.User{}, "", err
	}

	return *user, token, nil
}
