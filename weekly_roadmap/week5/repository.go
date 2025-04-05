package week5

import (
    "gorm.io/gorm"
)

type StudentRepository struct {
    db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
    return &StudentRepository{db: db}
}

// Lấy danh sách sinh viên
func (r *StudentRepository) GetAll() ([]Student, error) {
    var students []Student
    if err := r.db.Find(&students).Error; err != nil {
        return nil, err
    }
    return students, nil
}

// Lấy thông tin sinh viên theo ID
func (r *StudentRepository) GetByID(id int) (*Student, error) {
    var student Student
    if err := r.db.First(&student, id).Error; err != nil {
        return nil, err
    }
    return &student, nil
}

// Thêm sinh viên mới
func (r *StudentRepository) Add(student *Student) error {
    return r.db.Create(student).Error
}

// Cập nhật thông tin sinh viên
func (r *StudentRepository) Update(student *Student) error {
    return r.db.Save(student).Error
}

// Xóa sinh viên
func (r *StudentRepository) Delete(id int) error {
    return r.db.Delete(&Student{}, id).Error
}