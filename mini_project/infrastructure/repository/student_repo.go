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

func (r *StudentRepository) GetAll() ([]models.Student, error) {
	var students []models.Student
	if err := r.db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (r *StudentRepository) Update(student *models.Student) error {
	return r.db.Save(student).Error
}

func (r *StudentRepository) Delete(id int) error {
	return r.db.Delete(&models.Student{}, id).Error
}
