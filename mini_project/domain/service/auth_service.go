package service

import (
	"context"
	"errors"
	"time"

	"modern-tech/mini_project/domain/models"
	"modern-tech/mini_project/infrastructure/redis"
	"modern-tech/mini_project/infrastructure/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("your_secret_key")

type AuthService struct {
	repo        *repository.StudentRepository
	redisClient *redis.RedisClient
}

func NewAuthService(repo *repository.StudentRepository, redisClient *redis.RedisClient) *AuthService {
	return &AuthService{repo: repo, redisClient: redisClient}
}

func (s *AuthService) Register(student *models.Student) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(student.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	student.Password = string(hashedPassword)
	return s.repo.CreateStudent(student)
}

func (s *AuthService) Login(email, password string) (string, error) {
	student, err := s.repo.GetStudentByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"student_id": student.ID,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	// Define a context for Redis operations
	ctx := context.Background()

	// Save token into Redis
	err = s.redisClient.Client.Set(ctx, "token:"+tokenString, student.ID, time.Hour*24).Err()
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
}

func (s *AuthService) Logout(tokenString string) error {
	// Define a context for Redis operations
	ctx := context.Background()

	return s.redisClient.Client.Del(ctx, "token:"+tokenString).Err()
}
