package usecase

import (
	"context"
	"time"

	"modern-tech/mini_project/internal/domain"
	"modern-tech/mini_project/internal/repository"
	"modern-tech/mini_project/pkg/logger"
)

// UserUsecase struct
type UserUsecase struct {
	UserRepository repository.UserRepository
}

// NewUserUsecase function
func NewUserUsecase(userRepository repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		UserRepository: userRepository,
	}
}

// CreateUser function
func (u *UserUsecase) CreateUser(ctx context.Context, user domain.User) error {
	user.Password = util.HashPassword(user.Password)
	user.CreatedAt = time.Now().Unix()
	err := u.UserRepository.CreateUser(ctx, user)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

// GetUserByUsername function
func (u *UserUsecase) GetUserByUsername(ctx context.Context, username string) (domain.User, error) {
	user, err := u.UserRepository.GetUserByUsername(ctx, username)
	if err != nil {
		logger.Error(err.Error())
		return user, err
	}
	return user, nil
}

// GetUserByID function
func (u *UserUsecase) GetUserByID(ctx context.Context, id string) (domain.User, error) {
	user, err := u.UserRepository.GetUserByID(ctx, id)
	if err != nil {
		logger.Error(err.Error())
		return user, err
	}
	return user, nil
}
