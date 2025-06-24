package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Username     string             `bson:"username"`
	Email        string             `bson:"email"`
	PasswordHash string             `bson:"password_hash"`
	Role         string             `bson:"role"`
	IsVerified   bool               `bson:"is_verified"`
	CreatedAt    time.Time          `bson:"created_at"`
}
