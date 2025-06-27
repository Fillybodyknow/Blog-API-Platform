package router

import (
	"github.com/Fillybodyknow/blog-api/internal/handler"
	"github.com/Fillybodyknow/blog-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

type LikeRouter struct {
	LikeHandler *handler.LikeHandler
}

func NewLikeRouter(likeHandler *handler.LikeHandler) *LikeRouter {
	return &LikeRouter{LikeHandler: likeHandler}
}

func (r *LikeRouter) LikeRouters(rg *gin.RouterGroup) {
	rg.Use(middleware.AuthMiddleware())
	rg.POST("/:post_id/like", r.LikeHandler.LikePost)
	rg.DELETE("/:post_id/like", r.LikeHandler.UnlikePost)
}
