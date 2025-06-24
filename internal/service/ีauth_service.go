package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Fillybodyknow/blog-api/internal/models"
	"github.com/Fillybodyknow/blog-api/internal/repository"
	"github.com/Fillybodyknow/blog-api/pkg/utility"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type OTPVerify struct {
	Email     string
	OTP       string
	ExpiredAt time.Time
}

var StoreOTP []OTPVerify

type AuthServiceInterface interface {
	Register(ctx context.Context, user *models.User) error
	Login(ctx context.Context, username string, password string) (models.User, string, error)
	SendOTP(UserID primitive.ObjectID, ctx context.Context) error
	VerifyOTP(UserID primitive.ObjectID, otp string) error
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

func (s *AuthService) SendOTP(UserID primitive.ObjectID, ctx context.Context) error {
	FoundUser, err := s.AuthRepository.FindUserByID(ctx, UserID)
	if err != nil || FoundUser == nil {
		return errors.New("ไม่พบบัญชีผู้ใช้")
	}
	var newStore []OTPVerify
	for _, v := range StoreOTP {
		if v.Email == FoundUser.Email && v.ExpiredAt.After(time.Now().Add(-1*time.Minute)) {
			return errors.New("กรุณารอสักครู่ก่อนขอ OTP ใหม่")
		}
		if v.Email != FoundUser.Email {
			newStore = append(newStore, v)
		}
	}
	StoreOTP = newStore

	OTP := utility.GenerateOTP()

	err = utility.SendEmail(FoundUser.Email, "OTP สำหรับยืนยันตัวตน", fmt.Sprintf("รหัส OTP คือ %s", OTP))
	if err != nil {
		return errors.New("ไม่สามารถส่ง OTP ได้")
	}

	StoreOTP = append(StoreOTP, OTPVerify{
		Email:     FoundUser.Email,
		OTP:       OTP,
		ExpiredAt: time.Now().Add(5 * time.Minute),
	})

	return nil
}

func (s *AuthService) VerifyOTP(userID primitive.ObjectID, otp string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := s.AuthRepository.FindUserByID(ctx, userID)
	if err != nil || user == nil {
		return errors.New("ไม่สามารถยืนยัน OTP ได้")
	}

	found := false
	var newStore []OTPVerify

	for _, v := range StoreOTP {

		if v.Email == user.Email && v.OTP == otp && v.ExpiredAt.After(time.Now()) {

			err := s.AuthRepository.UpdateVerifyUser(ctx, userID)
			if err != nil {
				return errors.New("ไม่สามารถอัปเดตสถานะยืนยันบัญชีได้")
			}
			found = true
			continue
		}
		newStore = append(newStore, v)
	}

	StoreOTP = newStore
	if !found {
		return errors.New("OTP ไม่ถูกต้องหรือหมดอายุแล้ว")
	}
	return nil
}
