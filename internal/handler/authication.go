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

// RegisterUser godoc
// @Summary สมัครสมาชิก
// @Description สร้างผู้ใช้งานใหม่ด้วย Username, Email และ Password
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.RegisterInput true "ข้อมูลผู้ใช้ใหม่"
// @Router /auth/register [post]
func (h *AuthHandler) RegisterUser(c *gin.Context) {

	var input models.RegisterInput

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

// RegisterUser godoc
// @Summary เข้าสู่ระบบ
// @Description เข้าสู่ระบบด้วย Username,Email และ Password
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.LoginUserInput true "ข้อมูลผู้ใช้"
// @Router /auth/login [post]
func (h *AuthHandler) LoginUser(c *gin.Context) {
	var input models.LoginUserInput

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

// SendOTP godoc
// @Summary ส่ง OTP
// @Description ส่ง OTP ไปยัง Email ของผู้ใช้
// @Tags Auth
// @Security BearerAuth
// @Accept json
// @Produce json
// @Router /auth/verify_otp [get]
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

// VerifyOTP godoc
// @Summary ยืนยัน OTP
// @Description ยืนยัน OTP ของผู้ใช้
// @Tags Auth
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param otp body models.VerifyOTPInput true "ข้อมูล OTP"
// @Router /auth/verify-otp [post]
func (h *AuthHandler) VerifyOTP(c *gin.Context) {
	UserIDStr, _ := c.Get("user_id")
	objID, err := primitive.ObjectIDFromHex(UserIDStr.(string))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var input models.VerifyOTPInput
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
