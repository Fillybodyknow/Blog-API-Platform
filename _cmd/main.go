package main

import (
	"log"
	"os"

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

	r := gin.Default()

	api := r.Group("/api")

	authGroup := api.Group("/auth")
	authRouter.AuthRoutes(authGroup)
	authRouter.OTPRoutes(authGroup)

	postGroup := api.Group("/posts")
	postRouter.PostRoutes(postGroup)

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
