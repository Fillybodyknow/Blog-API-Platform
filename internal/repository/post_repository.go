package repository

import (
	"context"

	"github.com/Fillybodyknow/blog-api/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepositoryInterface interface {
	InsertPost(ctx context.Context, post *models.Post) error
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
