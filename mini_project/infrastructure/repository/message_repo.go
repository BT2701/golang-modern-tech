package repository

import (
	"modern-tech/mini_project/domain/models"

	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) SaveMessage(message *models.Message) error {
	return r.db.Create(message).Error
}

func (r *MessageRepository) GetMessagesByReceiverID(receiverID int) ([]models.Message, error) {
	var messages []models.Message
	if err := r.db.Where("receiver_id = ?", receiverID).Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}
