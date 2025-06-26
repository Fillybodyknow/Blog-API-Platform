package handler

import (
	"strings"
	"time"

	"github.com/Fillybodyknow/blog-api/internal/models"
	"github.com/Fillybodyknow/blog-api/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostHandler struct {
	PostServiceInterface service.PostServiceInterface
}

func NewPostHandler(postService service.PostServiceInterface) *PostHandler {
	return &PostHandler{PostServiceInterface: postService}
}

func (h *PostHandler) CreatePost(c *gin.Context) {

	var input models.PostInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	UserIDstr, _ := c.Get("user_id")
	UserID, _ := primitive.ObjectIDFromHex(UserIDstr.(string))

	var Tags = strings.Split(input.Tags, ",")

	Post := models.Post{
		Title:     input.Title,
		Content:   input.Content,
		AuthorID:  UserID,
		Tags:      Tags,
		Published: true,
		Comments:  []models.Comment{},
		Likes:     []models.Like{},
		CreatedAt: time.Now(),
	}

	err := h.PostServiceInterface.CreatePost(&Post, c.GetString("role"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Post created successfully"})
}

func (h *PostHandler) GetAllPosts(c *gin.Context) {

	posts, err := h.PostServiceInterface.GetAllPosts()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"posts": posts})
}

func (h *PostHandler) GetMePosts(c *gin.Context) {

	UserIDstr, _ := c.Get("user_id")
	UserID, _ := primitive.ObjectIDFromHex(UserIDstr.(string))

	posts, err := h.PostServiceInterface.GetAuthorPosts(UserID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"posts": posts})
}

func (h *PostHandler) GetPostsFromTags(c *gin.Context) {

	tags := c.Query("tags")

	tagSplit := strings.Split(tags, ",")

	posts, err := h.PostServiceInterface.GetPostsFromTags(tagSplit)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"posts": posts})
}
