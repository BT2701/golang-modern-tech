package service

import (
	"errors"
	"modern-tech/mini_project/domain/models"
	"modern-tech/mini_project/infrastructure/repository"
	"modern-tech/mini_project/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
)

type StudentService struct {
	repo *repository.StudentRepository
}

func NewStudentService(repo *repository.StudentRepository) *StudentService {
	return &StudentService{repo: repo}
}

func (s *StudentService) Register(student *models.Student) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(student.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	student.Password = string(hashedPassword)
	return s.repo.CreateStudent(student)
}

func (s *StudentService) Login(email, password string) (string, error) {
	student, err := s.repo.GetStudentByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	return jwt.GenerateToken(student.ID)
}
