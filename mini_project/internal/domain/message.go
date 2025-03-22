package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    SenderID  primitive.ObjectID `bson:"sender_id"`
    ReceiverID primitive.ObjectID `bson:"receiver_id"`
    Content   string             `bson:"content"`
    CreatedAt int64              `bson:"created_at"`
}
