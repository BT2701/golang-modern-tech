package service

import (
	"modern-tech/mini_project/domain/models"
	"modern-tech/mini_project/infrastructure/repository"
)

type MessageService struct {
	repo *repository.MessageRepository
}

func NewMessageService(repo *repository.MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) SaveMessage(message *models.Message) error {
	return s.repo.SaveMessage(message)
}

func (s *MessageService) GetMessagesByReceiverID(receiverID int) ([]models.Message, error) {
	return s.repo.GetMessagesByReceiverID(receiverID)
}
