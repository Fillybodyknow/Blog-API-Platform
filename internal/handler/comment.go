package handler

import (
	"github.com/Fillybodyknow/blog-api/internal/models"
	"github.com/Fillybodyknow/blog-api/internal/service"
	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	CommentServiceInterface service.CommentServiceInterface
}

func NewCommentHandler(commentService service.CommentServiceInterface) *CommentHandler {
	return &CommentHandler{CommentServiceInterface: commentService}
}

func (h *CommentHandler) Comment(c *gin.Context) {
	var input models.CommentInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	UserID, _ := c.Get("user_id")
	PostID := c.Param("post_id")

	err := h.CommentServiceInterface.Comment(input.Content, PostID, UserID.(string))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Comment created successfully"})

}

func (h *CommentHandler) EditComment(c *gin.Context) {
	var input models.CommentInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	PostID := c.Param("post_id")
	CommentID := c.Param("comment_id")
	UserID, _ := c.Get("user_id")

	err := h.CommentServiceInterface.EditComment(input.Content, PostID, CommentID, UserID.(string))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Comment edited successfully"})
}

func (h *CommentHandler) DeleteComment(c *gin.Context) {
	PostID := c.Param("post_id")
	CommentID := c.Param("comment_id")
	UserID, _ := c.Get("user_id")

	err := h.CommentServiceInterface.DeleteComment(PostID, CommentID, UserID.(string))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Comment deleted successfully"})
}
