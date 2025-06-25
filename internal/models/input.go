package models

type PostInput struct {
	Title   string `json:"title" form:"title" validate:"required"`
	Content string `json:"content" form:"content" validate:"required"`
	Tags    string `json:"tags" form:"tags" validate:"required"`
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

type VerifyOTPInput struct {
	OTP string `json:"otp" form:"otp" binding:"required"`
}
