package repository

import (
	"modern-tech/mini_project/domain/models"

	"gorm.io/gorm"
)

type StudentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

func (r *StudentRepository) CreateStudent(student *models.Student) error {
	return r.db.Create(student).Error
}

func (r *StudentRepository) GetStudentByEmail(email string) (*models.Student, error) {
	var student models.Student
	if err := r.db.Where("email = ?", email).First(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}
