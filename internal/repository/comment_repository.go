package repository

import (
	"context"

	"github.com/Fillybodyknow/blog-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentRepositoryInterface interface {
	InsertComment(ctx context.Context, comment *models.Comment, PostID primitive.ObjectID) error
	GetComments(ctx context.Context, PostID primitive.ObjectID) ([]models.Comment, error)
	UpdateComment(ctx context.Context, content string, PostID primitive.ObjectID, CommentID primitive.ObjectID) error
	DeleteComment(ctx context.Context, PostID, CommentID primitive.ObjectID) error
}

func (r *PostRepository) InsertComment(ctx context.Context, comment *models.Comment, PostID primitive.ObjectID) error {
	err := r.Collection.FindOneAndUpdate(ctx, bson.M{"_id": PostID}, bson.M{"$push": bson.M{"comment": comment}}).Err()
	return err
}

func (r *PostRepository) GetComments(ctx context.Context, PostID primitive.ObjectID) ([]models.Comment, error) {
	var post models.Post
	err := r.Collection.FindOne(ctx, bson.M{"_id": PostID}).Decode(&post)
	if err != nil {
		return nil, err
	}
	return post.Comments, nil
}

func (r *PostRepository) UpdateComment(ctx context.Context, content string, PostID primitive.ObjectID, CommentID primitive.ObjectID) error {
	filter := bson.M{
		"_id":         PostID,
		"comment._id": CommentID,
	}
	update := bson.M{
		"$set": bson.M{
			"comment.$.content": content,
		},
	}
	_, err := r.Collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *PostRepository) DeleteComment(ctx context.Context, PostID, CommentID primitive.ObjectID) error {
	filter := bson.M{"_id": PostID}
	update := bson.M{"$pull": bson.M{"comment": bson.M{"_id": CommentID}}}
	_, err := r.Collection.UpdateOne(ctx, filter, update)
	return err
}
