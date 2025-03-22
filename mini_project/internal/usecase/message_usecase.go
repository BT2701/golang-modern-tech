package usecase

import (
	"context"
	"time"

	"modern-tech/mini_project/internal/domain"
	"modern-tech/mini_project/internal/repository"
	"modern-tech/mini_project/pkg/logger"
)

// MessageUsecase struct
type MessageUsecase struct {
	MessageRepository repository.MessageRepository
}

// NewMessageUsecase function
func NewMessageUsecase(messageRepository repository.MessageRepository) *MessageUsecase {
	return &MessageUsecase{
		MessageRepository: messageRepository,
	}
}

// CreateMessage function
func (u *MessageUsecase) CreateMessage(ctx context.Context, message domain.Message) error {
	message.CreatedAt = time.Now().Unix()
	err := u.MessageRepository.CreateMessage(ctx, message)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

// GetMessagesByUserID function
func (u *MessageUsecase) GetMessagesByUserID(ctx context.Context, userID string) ([]domain.Message, error) {
	messages, err := u.MessageRepository.GetMessagesByUserID(ctx, userID)
	if err != nil {
		logger.Error(err.Error())
		return messages, err
	}
	return messages, nil
}

// GetMessageByID function
func (u *MessageUsecase) GetMessageByID(ctx context.Context, id string) (domain.Message, error) {
	message, err := u.MessageRepository.GetMessageByID(ctx, id)
	if err != nil {
		logger.Error(err.Error())
		return message, err
	}
	return message, nil
}

// DeleteMessage function
func (u *MessageUsecase) DeleteMessage(ctx context.Context, id string) error {
	err := u.MessageRepository.DeleteMessage(ctx, id)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

// UpdateMessage function
func (u *MessageUsecase) UpdateMessage(ctx context.Context, id string, message domain.Message) error {
	err := u.MessageRepository.UpdateMessage(ctx, id, message)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

// GetMessageBySenderAndReceiverID function
func (u *MessageUsecase) GetMessageBySenderAndReceiverID(ctx context.Context, senderID, receiverID string) ([]domain.Message, error) {
	messages, err := u.MessageRepository.GetMessageBySenderAndReceiverID(ctx, senderID, receiverID)
	if err != nil {
		logger.Error(err.Error())
		return messages, err
	}
	return messages, nil
}

// GetMessageBySenderID function
func (u *MessageUsecase) GetMessageBySenderID(ctx context.Context, senderID string) ([]domain.Message, error) {
	messages, err := u.MessageRepository.GetMessageBySenderID(ctx, senderID)
	if err != nil {
		logger.Error(err.Error())
		return messages, err
	}
	return messages, nil
}

// GetMessageByReceiverID function
func (u *MessageUsecase) GetMessageByReceiverID(ctx context.Context, receiverID string) ([]domain.Message, error) {
	messages, err := u.MessageRepository.GetMessageByReceiverID(ctx, receiverID)
	if err != nil {
		logger.Error(err.Error())
		return messages, err
	}
	return messages, nil
}
