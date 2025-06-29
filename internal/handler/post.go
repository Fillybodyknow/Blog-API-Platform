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

// @Summary สร้าง Post
// @Description ต้องยืนยันตัวตนก่อนจึงจะสามารถสร้าง Post ได้
// @Tags Post
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Security BearerAuth
// @Param title formData string true "Title" example("Hello World")
// @Param content formData string true "Content" example("Hello World Content")
// @Param tags formData string true "Tags" example("tag1,tag2,tag3")
// @Router /posts/create [post]
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

// @Summary แสดง Post ทั้งหมด
// @Description แสดง Post ทั้งหมด ไม่ต้องยืนยันตัวตน
// @Tags Post
// @Produce json
// @Router /posts/ [get]
func (h *PostHandler) GetAllPosts(c *gin.Context) {

	posts, err := h.PostServiceInterface.GetAllPosts()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"posts": posts})
}

// @Summary แสดง Post ของตัวเอง
// @Description แสดง Post ของตัวเอง ไม่ต้องยืนยันตัวตน
// @Tags Post
// @Produce json
// @Security BearerAuth
// @Router /posts/me [get]
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

// @Summary แสดง Post ตาม Tags
// @Description แสดง Post ตาม Tags ไม่ต้องยืนยันตัวตน
// @Tags Post
// @Produce json
// @Param tags query string true "Tags" example("tag1,tag2,tag3")
// @Router /posts [get]
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

// @Summary แสดง Post ตาม ID
// @Description แสดง Post ตาม ID ไม่ต้องยืนยันตัวตน
// @Tags Post
// @Produce json
// @Param post_id path string true "Post ID"
// @Router /posts/{post_id} [get]
func (h *PostHandler) GetPostByID(c *gin.Context) {

	idStr := c.Param("post_id")

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	post, err := h.PostServiceInterface.GetPostByID(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"post": post})
}

// @Summary แก้ไข Post ของตัวเอง
// @Description แก้ไข Post ของตัวเอง ต้องยืนยันตัวตนก่อนจึงจะสามารถแก้ไข Post ได้
// @Tags Post
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Security BearerAuth
// @Param post_id path string true "Post ID"
// @Param title formData string true "Title" example("Hello World")
// @Param content formData string true "Content" example("Hello World")
// @Param tags formData string true "Tags" example("tag1,tag2,tag3")
// @Router /posts/{post_id} [put]
func (h *PostHandler) EditPost(c *gin.Context) {

	var input models.PostInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	idStr := c.Param("post_id")
	UserID, _ := c.Get("user_id")
	Role, _ := c.Get("role")

	EditForm := models.Post{
		Title:   input.Title,
		Content: input.Content,
		Tags:    strings.Split(input.Tags, ","),
	}

	err := h.PostServiceInterface.EditMePost(&EditForm, Role.(string), UserID.(string), idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Post edited successfully"})
}

// @Summary ลบ Post ของตัวเอง
// @Description ลบ Post ของตัวเอง ต้องยืนยันตัวตนก่อนจึงจะสามารถลบ Post ได้
// @Tags Post
// @Security BearerAuth
// @Param post_id path string true "Post ID"
// @Router /posts/{post_id} [delete]
func (h *PostHandler) DeletePost(c *gin.Context) {

	idStr := c.Param("post_id")
	UserID, _ := c.Get("user_id")
	Role, _ := c.Get("role")

	err := h.PostServiceInterface.DeletePostByID(idStr, UserID.(string), Role.(string))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Post deleted successfully"})
}
