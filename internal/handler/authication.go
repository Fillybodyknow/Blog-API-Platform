package handler

import (
	"time"

	"github.com/Fillybodyknow/blog-api/internal/models"
	"github.com/Fillybodyknow/blog-api/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthHandler struct {
	AuthServiceInterface service.AuthServiceInterface
}

func NewAuthHandler(authService service.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{AuthServiceInterface: authService}
}

type RegisterInput struct {
	Username string `json:"username" form:"username" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type LoginUserInput struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (h *AuthHandler) RegisterUser(c *gin.Context) {

	var input RegisterInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username:     input.Username,
		Email:        input.Email,
		PasswordHash: input.Password,
		Role:         "user",
		IsVerified:   false,
		CreatedAt:    time.Now(),
	}

	err := h.AuthServiceInterface.Register(c, &user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Registered successfully"})

}

func (h *AuthHandler) LoginUser(c *gin.Context) {
	var input LoginUserInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	user, token, err := h.AuthServiceInterface.Login(c, input.Username, input.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"user_id": user.ID, "role": user.Role, "token": token})
}

func (h *AuthHandler) OTP(c *gin.Context) {
	UserIDStr, _ := c.Get("user_id")
	objID, err := primitive.ObjectIDFromHex(UserIDStr.(string))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = h.AuthServiceInterface.SendOTP(objID, c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "ส่ง OTP สําเร็จ"})
}

type VerifyOTPInput struct {
	OTP string `json:"otp" form:"otp" binding:"required"`
}

func (h *AuthHandler) VerifyOTP(c *gin.Context) {
	UserIDStr, _ := c.Get("user_id")
	objID, err := primitive.ObjectIDFromHex(UserIDStr.(string))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var input VerifyOTPInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = h.AuthServiceInterface.VerifyOTP(objID, input.OTP)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "ยืนยัน OTP สําเร็จ"})
}
