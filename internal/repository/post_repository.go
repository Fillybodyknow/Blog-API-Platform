package repository

import (
	"context"
	"strings"

	"github.com/Fillybodyknow/blog-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepositoryInterface interface {
	Insert(ctx context.Context, post *models.Post) error
	Get(ctx context.Context) ([]models.Post, error)
	FindByAuthorID(ctx context.Context, AuthorId primitive.ObjectID) ([]models.Post, error)
	FindByTags(ctx context.Context, tags []string) ([]models.Post, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*models.Post, error)
	Update(ctx context.Context, id primitive.ObjectID, post *models.Post) error
}

type PostRepository struct {
	Collection *mongo.Collection
}

func NewPostRepository(collection *mongo.Collection) *PostRepository {
	return &PostRepository{
		Collection: collection,
	}
}

func (r *PostRepository) Insert(ctx context.Context, post *models.Post) error {
	_, err := r.Collection.InsertOne(ctx, post)
	return err
}

func (r *PostRepository) Get(ctx context.Context) ([]models.Post, error) {

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

func (r *PostRepository) FindByAuthorID(ctx context.Context, authorID primitive.ObjectID) ([]models.Post, error) {
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

func (r *PostRepository) FindByTags(ctx context.Context, tags []string) ([]models.Post, error) {

	var lowerTags []string
	for _, tag := range tags {
		lowerTags = append(lowerTags, strings.ToLower(tag))
	}

	filter := bson.M{
		"$expr": bson.M{
			"$setIsSubset": []interface{}{
				lowerTags,
				bson.M{
					"$map": bson.M{
						"input": "$tags",
						"as":    "t",
						"in":    bson.M{"$toLower": "$$t"},
					},
				},
			},
		},
	}

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

func (r *PostRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Post, error) {
	var post models.Post

	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&post)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}
	return &post, nil

}

func (r *PostRepository) Update(ctx context.Context, id primitive.ObjectID, post *models.Post) error {
	update := bson.M{
		"$set": bson.M{
			"title":   post.Title,
			"content": post.Content,
			"tags":    post.Tags,
		},
	}
	_, err := r.Collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	return err
}
