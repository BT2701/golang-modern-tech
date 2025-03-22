package repository

import (
	"context"
	"fmt"
	"modern-tech/mini_project/internal/domain"
	"modern-tech/mini_project/pkg/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MessageRepository struct
type MessageRepository struct {
	Collection *mongo.Collection
}

// NewMessageRepository function
func NewMessageRepository(db *mongo.Database) *MessageRepository {
	return &MessageRepository{
		Collection: db.Collection("messages"),
	}
}

// CreateMessage function
func (r *MessageRepository) CreateMessage(ctx context.Context, message domain.Message) error {
	message.CreatedAt = time.Now().Unix()
	_, err := r.Collection.InsertOne(ctx, message)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

// GetMessagesByUserID function
func (r *MessageRepository) GetMessagesByUserID(ctx context.Context, userID string) ([]domain.Message, error) {
	var messages []domain.Message
	cursor, err := r.Collection.Find(ctx, bson.M{
		"$or": []bson.M{
			{"sender_id": userID},
			{"receiver_id": userID},
		},
	}, options.Find().SetSort(bson.M{"created_at": -1}))
	if err != nil {
		logger.Error(err.Error())
		return messages, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var message domain.Message
		err := cursor.Decode(&message)
		if err != nil {
			logger.Error(err.Error())
			return messages, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}

// GetMessageByID function
func (r *MessageRepository) GetMessageByID(ctx context.Context, id string) (domain.Message, error) {
	var message domain.Message
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error(err.Error())
		return message, err
	}
	err = r.Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&message)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return message, fmt.Errorf("message not found")
		}
		logger.Error(err.Error())
		return message, err
	}
	return message, nil
}

// UpdateMessage function
func (r *MessageRepository) UpdateMessage(ctx context.Context, id string, message domain.Message) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	_, err = r.Collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": message})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

// DeleteMessage function
func (r *MessageRepository) DeleteMessage(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	_, err = r.Collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
