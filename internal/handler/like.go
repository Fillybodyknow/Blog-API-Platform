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
