package models

type PostInput struct {
	Title   string `json:"title" form:"title" validate:"required"`
	Content string `json:"content" form:"content" validate:"required"`
	Tags    string `json:"tags" form:"tags" validate:"required"`
}

type RegisterInput struct {
	Username string `json:"username" form:"username" binding:"required" example:"blog"`
	Email    string `json:"email" form:"email" binding:"required,email" example:"blogg@example.com"`
	Password string `json:"password" form:"password" binding:"required" example:"Strong@Password123"`
}

type LoginUserInput struct {
	Username string `json:"username" form:"username" binding:"required" example:"blog or blogg@example.com"`
	Password string `json:"password" form:"password" binding:"required" example:"Strong@Password123"`
}

type VerifyOTPInput struct {
	OTP string `json:"otp" form:"otp" binding:"required" example:"123456"`
}

type CommentInput struct {
	Content string `json:"content" form:"content" validate:"required"`
}
