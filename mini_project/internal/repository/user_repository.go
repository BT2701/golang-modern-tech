package repository

import (
	"context"
	"fmt"
	"modern-tech/mini_project/internal/domain"
	"modern-tech/mini_project/pkg/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UserRepository struct
type UserRepository struct {
	Collection *mongo.Collection
}

// NewUserRepository function
func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		Collection: db.Collection("users"),
	}
}

// CreateUser function
func (r *UserRepository) CreateUser(ctx context.Context, user domain.User) error {
	user.CreatedAt = time.Now().Unix()
	_, err := r.Collection.InsertOne(ctx, user)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

// GetUserByUsername function
func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (domain.User, error) {
	var user domain.User
	err := r.Collection.FindOne(ctx, bson.M
		{"username": username},
	).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, fmt.Errorf("user not found")
		}
		logger.Error(err.Error())
		return user
	}
	return user, nil
}

// GetUserByID function
func (r *UserRepository) GetUserByID(ctx context.Context, id string) (domain.User, error) {
	var user domain.User
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error(err.Error())
		return user, err
	}
	err = r.Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, fmt.Errorf("user not found")
		}
		logger.Error(err.Error())
		return user, err
	}
	return user, nil
}

// UpdateUser function
func (r *UserRepository) UpdateUser(ctx context.Context, id string, user domain.User) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	_, err = r.Collection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": user})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

// DeleteUser function
func (r *UserRepository) DeleteUser(ctx context.Context, id string) error {
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

// GetAllUsers function
func (r *UserRepository) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		logger.Error(err.Error())
		return users, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &users)
	if err != nil {
		logger.Error(err.Error())
		return users, err
	}
	return users, nil
}

// GetUserCount function
func (r *UserRepository) GetUserCount(ctx context.Context) (int64, error) {
	count, err := r.Collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		logger.Error(err.Error())
		return 0, err
	}
	return count, nil
}

// GetUserByUsernameAndPassword function
func (r *UserRepository) GetUserByUsernameAndPassword(ctx context.Context, username, password string) (domain.User, error) {
	var user domain.User
	err := r.Collection.FindOne(ctx, bson.M
		{"username": username, "password": password},
	).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, fmt.Errorf("user not found")
		}
		logger.Error(err.Error())
		return user
	}
	return user, nil
}
