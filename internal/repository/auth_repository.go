package repository

import (
	"context"

	"github.com/Fillybodyknow/blog-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepositoryInterface interface {
	InsertUser(ctx context.Context, user *models.User) error
	FindByEmailOrUsername(ctx context.Context, emailorusername string) (*models.User, error)
	FindUserByID(ctx context.Context, id primitive.ObjectID) (*models.User, error)
	UpdateVerifyUser(ctx context.Context, id primitive.ObjectID) error
}

type AuthRepository struct {
	Collection *mongo.Collection
}

func NewAuthRepository(collection *mongo.Collection) *AuthRepository {
	return &AuthRepository{
		Collection: collection,
	}
}

func (r *AuthRepository) InsertUser(ctx context.Context, user *models.User) error {
	_, err := r.Collection.InsertOne(ctx, user)
	return err
}

func (r *AuthRepository) FindByEmailOrUsername(ctx context.Context, emailorusername string) (*models.User, error) {
	var user models.User
	err := r.Collection.FindOne(ctx, bson.M{"$or": []bson.M{{"email": emailorusername}, {"username": emailorusername}}}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *AuthRepository) FindUserByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	var user models.User
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepository) UpdateVerifyUser(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.Collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"is_verified": true, "role": "editor"}})
	return err
}
