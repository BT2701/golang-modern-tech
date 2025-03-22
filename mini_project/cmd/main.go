package cmd

import (
	"context"
	"fmt"
	"os"
	"modern-tech/mini_project/config"
	"modern-tech/mini_project/internal/repository"

	"modern-tech/mini_project/internal/socket"
	"modern-tech/mini_project/internal/usecase"
	"modern-tech/mini_project/pkg/logger"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// Init logger
	logger.InitLogger(cfg.LogLevel)

	// Init repository
	repo, err := repository.NewRepository(cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// Init usecase
	userUsecase := usecase.NewUserUsecase(repo.UserRepository)
	messageUsecase := usecase.NewMessageUsecase(repo.MessageRepository)

	// Init socket
	messageSocket := socket.NewMessageSocket(messageUsecase)

	// Start socket
	go messageSocket.Start()

	// Create user
	user := domain.User{
		Username: "john_doe",
		Password: "password"
	}
	userUsecase.CreateUser(context.Background(), user)

	// Get user by username
	user, err := userUsecase.GetUserByUsername(context.Background(), "john_doe")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	fmt.Println(user)

	// Get user by ID
	user, err = userUsecase.GetUserByID(context.Background(), user.ID)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// Create message
	message := domain.Message{
		Content: "Hello, world!",
	}

	messageUsecase.CreateMessage(context.Background(), message)
}
