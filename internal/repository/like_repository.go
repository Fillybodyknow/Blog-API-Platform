package repository

import (
	"context"

	"github.com/Fillybodyknow/blog-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LikeRepositoryInterface interface {
	InsertLike(ctx context.Context, like *models.Like, PostID primitive.ObjectID) error
	DeleteLike(ctx context.Context, PostID primitive.ObjectID, LikeID primitive.ObjectID) error
	GetLikes(ctx context.Context, PostID primitive.ObjectID) ([]models.Like, error)
	FindLikeByUserID(ctx context.Context, PostID primitive.ObjectID, userID primitive.ObjectID) (*models.Like, error)
	FindLikeByID(ctx context.Context, PostID primitive.ObjectID, LikeID primitive.ObjectID) (*models.Like, error)
}

func (r *PostRepository) InsertLike(ctx context.Context, like *models.Like, PostID primitive.ObjectID) error {
	err := r.Collection.FindOneAndUpdate(ctx, bson.M{"_id": PostID}, bson.M{"$push": bson.M{"likes": like}}).Err()
	return err
}

func (r *PostRepository) DeleteLike(ctx context.Context, PostID primitive.ObjectID, LikeID primitive.ObjectID) error {
	filter := bson.M{"_id": PostID}
	update := bson.M{"$pull": bson.M{"likes": bson.M{"_id": LikeID}}}
	_, err := r.Collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *PostRepository) GetLikes(ctx context.Context, PostID primitive.ObjectID) ([]models.Like, error) {
	var post models.Post
	err := r.Collection.FindOne(ctx, bson.M{"_id": PostID}).Decode(&post)
	if err != nil {
		return nil, err
	}
	return post.Likes, nil
}

func (r *PostRepository) FindLikeByUserID(ctx context.Context, PostID primitive.ObjectID, userID primitive.ObjectID) (*models.Like, error) {
	filter := bson.M{"_id": PostID, "likes.user_id": userID}
	projection := bson.M{"likes": bson.M{"$elemMatch": bson.M{"user_id": userID}}}

	var result struct {
		Likes []models.Like `bson:"likes"`
	}

	err := r.Collection.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		return nil, err
	}
	if len(result.Likes) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return &result.Likes[0], nil
}

func (r *PostRepository) FindLikeByID(ctx context.Context, PostID primitive.ObjectID, LikeID primitive.ObjectID) (*models.Like, error) {
	filter := bson.M{"_id": PostID, "likes._id": LikeID}
	projection := bson.M{"likes": bson.M{"$elemMatch": bson.M{"_id": LikeID}}}

	var result struct {
		Likes []models.Like `bson:"likes"`
	}

	err := r.Collection.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		return nil, err
	}
	if len(result.Likes) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return &result.Likes[0], nil
}
