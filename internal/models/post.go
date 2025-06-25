package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title"`
	Content   string             `bson:"content"`
	AuthorID  primitive.ObjectID `bson:"author_id"`
	Tags      []string           `bson:"tags"`
	Published bool               `bson:"published"`
	Comments  []Comment          `bson:"comment"`
	Likes     []Like             `bson:"likes"`
	CreatedAt time.Time          `bson:"created_at"`
}

type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id"`
	Content   string             `bson:"content"`
	CreatedAt time.Time          `bson:"created_at"`
}

type Like struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserID primitive.ObjectID `bson:"user_id"`
}

type Tag struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
}
