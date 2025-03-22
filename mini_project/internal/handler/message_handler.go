package handler

import (
	"modern-tech/mini_project/internal/domain"
	"modern-tech/mini_project/internal/usecase"
	"modern-tech/mini_project/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MessageHandler struct
type MessageHandler struct {
	MessageUsecase usecase.MessageUsecase
}

// NewMessageHandler function
func NewMessageHandler(messageUsecase usecase.MessageUsecase) *MessageHandler {
	return &MessageHandler{
		MessageUsecase: messageUsecase,
	}
}

// CreateMessage function
func (h *MessageHandler) CreateMessage(c *gin.Context) {
	var message domain.Message
	err := c.BindJSON(&message)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.MessageUsecase.CreateMessage(c, message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Message created"})
}

// GetMessagesByUserID function
func (h *MessageHandler) GetMessagesByUserID(c *gin.Context) {
	userID := c.Param("userID")

	messages, err := h.MessageUsecase.GetMessagesByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}

// GetMessageByID function
func (h *MessageHandler) GetMessageByID(c *gin.Context) {
	id := c.Param("id")

	message, err := h.MessageUsecase.GetMessageByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}

// DeleteMessage function
func (h *MessageHandler) DeleteMessage(c *gin.Context) {
	id := c.Param("id")

	err := h.MessageUsecase.DeleteMessage(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message deleted"})
}

// UpdateMessage function
func (h *MessageHandler) UpdateMessage(c *gin.Context) {
	id := c.Param("id")

	var message domain.Message
	err := c.BindJSON(&message)
	if err != nil {
		logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.MessageUsecase.UpdateMessage(c, id, message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message updated"})
}
