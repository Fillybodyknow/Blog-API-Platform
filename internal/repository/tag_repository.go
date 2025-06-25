package repository

import (
	"context"

	"github.com/Fillybodyknow/blog-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TagRepositoryInterface interface {
	InsertTag(ctx context.Context, tag *models.Tag) error
	FindTagByName(ctx context.Context, name string) (*models.Tag, error)
}

type TagRepository struct {
	Collection *mongo.Collection
}

func NewTagRepository(collection *mongo.Collection) *TagRepository {
	return &TagRepository{
		Collection: collection,
	}
}

func (r *TagRepository) InsertTag(ctx context.Context, tag *models.Tag) error {
	_, err := r.Collection.InsertOne(ctx, tag)
	return err
}

func (r *TagRepository) FindTagByName(ctx context.Context, name string) (*models.Tag, error) {
	var tag models.Tag

	err := r.Collection.FindOne(ctx, bson.M{
		"name": bson.M{
			"$regex":   name,
			"$options": "i",
		},
	}).Decode(&tag)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &tag, nil
}
