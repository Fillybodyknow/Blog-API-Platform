package repository

import (
	"context"

	"github.com/Fillybodyknow/blog-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepositoryInterface interface {
	InsertPost(ctx context.Context, post *models.Post) error
	GetPosts(ctx context.Context) ([]models.Post, error)
	FindPostByAuthorID(ctx context.Context, AuthorId primitive.ObjectID) ([]models.Post, error)
}

type PostRepository struct {
	Collection *mongo.Collection
}

func NewPostRepository(collection *mongo.Collection) *PostRepository {
	return &PostRepository{
		Collection: collection,
	}
}

func (r *PostRepository) InsertPost(ctx context.Context, post *models.Post) error {
	_, err := r.Collection.InsertOne(ctx, post)
	return err
}

func (r *PostRepository) GetPosts(ctx context.Context) ([]models.Post, error) {

	var posts []models.Post

	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostRepository) FindPostByAuthorID(ctx context.Context, authorID primitive.ObjectID) ([]models.Post, error) {
	filter := bson.M{"author_id": authorID}

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var posts []models.Post
	if err := cursor.All(ctx, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}
