package router

import (
	"github.com/Fillybodyknow/blog-api/internal/handler"
	"github.com/Fillybodyknow/blog-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

type CommentRouter struct {
	CommentHandler *handler.CommentHandler
}

func NewCommentRouter(commentHandler *handler.CommentHandler) *CommentRouter {
	return &CommentRouter{CommentHandler: commentHandler}
}

func (r *CommentRouter) CommentRoutes(rg *gin.RouterGroup) {
	rg.Use(middleware.AuthMiddleware())
	rg.POST("/:post_id/comment", r.CommentHandler.Comment)
	rg.PUT("/:post_id/comment/:comment_id", r.CommentHandler.EditComment)
}
