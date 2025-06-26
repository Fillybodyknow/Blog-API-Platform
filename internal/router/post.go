package router

import (
	"github.com/Fillybodyknow/blog-api/internal/handler"
	"github.com/Fillybodyknow/blog-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

type PostRouter struct {
	PostHandler *handler.PostHandler
}

func NewPostRouter(postHandler *handler.PostHandler) *PostRouter {
	return &PostRouter{PostHandler: postHandler}
}

func (r *PostRouter) PostRoutes(rg *gin.RouterGroup) {
	rg.GET("/", r.PostHandler.GetAllPosts)
	rg.GET("", r.PostHandler.GetPostsFromTags)
	rg.GET("/:id", r.PostHandler.GetPostByID)
}

func (r *PostRouter) PostMiddlewareRoutes(rg *gin.RouterGroup) {
	rg.Use(middleware.AuthMiddleware())
	rg.POST("/create", r.PostHandler.CreatePost)
	rg.GET("/me", r.PostHandler.GetMePosts)
	rg.PUT("/edit/:id", r.PostHandler.EditPost)
}
