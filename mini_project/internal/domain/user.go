package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Username  string             `bson:"username" json:"username"`
    Password  string             `bson:"password" json:"password"`
    CreatedAt int64              `bson:"created_at" json:"created_at"`
}
