package socket

import (
	"encoding/json"
	"fmt"
	"modern-tech/mini_project/internal/domain"
	"modern-tech/mini_project/internal/usecase"
	"modern-tech/mini_project/pkg/logger"
	"net"
	"os"
)

// MessageSocket struct
type MessageSocket struct {
	MessageUsecase usecase.MessageUsecase
}

// NewMessageSocket function
func NewMessageSocket(messageUsecase usecase.MessageUsecase) *MessageSocket {
	return &MessageSocket{
		MessageUsecase: messageUsecase,
	}
}

// Start function
func (s *MessageSocket) Start() {
	fmt.Println("Socket server started")
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			logger.Error(err.Error())
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *MessageSocket) handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			logger.Error(err.Error())
			break
		}
		var message domain.Message
		err = json.Unmarshal(buf[:n], &message)
		if err != nil {
			logger.Error(err.Error())
			break
		}
		message.CreatedAt = util.GetUnixTimestamp()
		err = s.MessageUsecase.CreateMessage(nil, message)
		if err != nil {
			logger.Error(err.Error())
			break
		}
	}
}
