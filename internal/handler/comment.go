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

// @Summary Create Comment
// @Description สร้าง Comment ด้วยการกรอก Content
// @Tags Comment
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Security BearerAuth
// @Param post_id path string true "Post ID" example("123456789012345678901234")
// @Param content formData string true "Content" example("Hello World")
// @Success 200 {object} map[string]string
// @Router /posts/{post_id}/comment [post]
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

// EditComment godoc
// @Summary Edit Comment
// @Description แก้ไข Comment ด้วนการกรอก Content
// @Tags Comment
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Security BearerAuth
// @Param post_id path string true "Post ID" example("123456789012345678901234")
// @Param comment_id path string true "Comment ID" example("123456789012345678901234")
// @Param content formData string true "Content" example("Hello World")
// @Router /posts/{post_id}/comment/{comment_id} [put]
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

// DeleteComment godoc
// @Summary Delete Comment
// @Description ลบ Comment ของตัวเอง
// @Tags Comment
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Security BearerAuth
// @Param post_id path string true "Post ID" example("123456789012345678901234")
// @Param comment_id path string true "Comment ID" example("123456789012345678901234")
// @Router /posts/{post_id}/comment/{comment_id} [delete]
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
