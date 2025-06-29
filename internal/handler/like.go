package handler

import (
	"github.com/Fillybodyknow/blog-api/internal/service"
	"github.com/gin-gonic/gin"
)

type LikeHandler struct {
	LikeServiceInterface service.LikeServiceInterface
}

func NewLikeHandler(likeService service.LikeServiceInterface) *LikeHandler {
	return &LikeHandler{LikeServiceInterface: likeService}
}

// @Summary Like
// @Description Like โพสต์
// @Tags Like
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Security BearerAuth
// @Param post_id path string true "Post ID"
// @Success 200 {object} map[string]string
// @Router /posts/{post_id}/like [post]
func (h *LikeHandler) LikePost(c *gin.Context) {
	UserID, _ := c.Get("user_id")
	PostID := c.Param("post_id")

	err := h.LikeServiceInterface.LikePost(PostID, UserID.(string))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Post liked successfully"})
}

// @Summary UnLike
// @Description UnLike โพสต์
// @Tags Like
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Security BearerAuth
// @Param post_id path string true "Post ID"
// @Success 200 {object} map[string]string
// @Router /posts/{post_id}/like [delete]
func (h *LikeHandler) UnlikePost(c *gin.Context) {
	UserID, _ := c.Get("user_id")
	PostID := c.Param("post_id")

	err := h.LikeServiceInterface.UnlikePost(PostID, UserID.(string))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Post unliked successfully"})
}
