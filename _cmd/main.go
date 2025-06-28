// @title           Blog API Platform
// @version         1.0
// @description     ระบบ Blog + Auth + OTP + Like/Comment ที่เขียนด้วย Go + Gin
// @contact.name    Filly
// @contact.email   fillybodyknow@gmail.com
// @license.name    MIT
// @host      localhost:8080
// @BasePath  /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	"log"
	"os"

	_ "github.com/Fillybodyknow/blog-api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/Fillybodyknow/blog-api/internal/config"
	"github.com/Fillybodyknow/blog-api/internal/handler"
	"github.com/Fillybodyknow/blog-api/internal/repository"
	"github.com/Fillybodyknow/blog-api/internal/router"
	"github.com/Fillybodyknow/blog-api/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := config.ConnectMongo()

	DBName := os.Getenv("DB_NAME")

	UserCollection := db.Database(DBName).Collection("users")
	authRepo := repository.NewAuthRepository(UserCollection)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)
	authRouter := router.NewAuthRouter(authHandler)

	PostCollection := db.Database(DBName).Collection("posts")
	TagCollection := db.Database(DBName).Collection("tags")
	tagRepo := repository.NewTagRepository(TagCollection)
	postRepo := repository.NewPostRepository(PostCollection)
	postService := service.NewPostService(postRepo, tagRepo)
	postHandler := handler.NewPostHandler(postService)
	postRouter := router.NewPostRouter(postHandler)

	CommentService := service.NewCommentService(postRepo)
	CommentHandler := handler.NewCommentHandler(CommentService)
	CommentRouter := router.NewCommentRouter(CommentHandler)

	LikeService := service.NewLikeService(postRepo)
	LikeHandler := handler.NewLikeHandler(LikeService)
	LikeRouter := router.NewLikeRouter(LikeHandler)

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")

	authGroup := api.Group("/auth")
	authRouter.AuthRoutes(authGroup)
	authRouter.OTPRoutes(authGroup)

	postGroup := api.Group("/posts")
	postRouter.PostRoutes(postGroup)
	postRouter.PostMiddlewareRoutes(postGroup)

	CommentRouter.CommentRoutes(postGroup)

	LikeRouter.LikeRouters(postGroup)

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
