package week7

import "gorm.io/gorm"

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) SaveMessage(message *Message) error {
	return r.db.Create(message).Error
}
