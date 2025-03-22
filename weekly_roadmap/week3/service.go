package week3

import (
    "errors"
    "sync"
)

type StudentService struct {
    mu       sync.Mutex
    students map[int]Student
    nextID   int
}

func NewStudentService() *StudentService {
    return &StudentService{
        students: make(map[int]Student),
        nextID:   1,
    }
}

func (s *StudentService) GetAll() []Student {
    s.mu.Lock()
    defer s.mu.Unlock()

    students := make([]Student, 0, len(s.students))
    for _, student := range s.students {
        students = append(students, student)
    }
    return students
}

func (s *StudentService) GetByID(id int) (Student, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    student, exists := s.students[id]
    if !exists {
        return Student{}, errors.New("student not found")
    }
    return student, nil
}

func (s *StudentService) Add(student Student) Student {
    s.mu.Lock()
    defer s.mu.Unlock()

    student.ID = s.nextID
    s.nextID++
    s.students[student.ID] = student
    return student
}

func (s *StudentService) Update(id int, student Student) (Student, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    _, exists := s.students[id]
    if !exists {
        return Student{}, errors.New("student not found")
    }

    student.ID = id
    s.students[id] = student
    return student, nil
}

func (s *StudentService) Delete(id int) error {
    s.mu.Lock()
    defer s.mu.Unlock()

    _, exists := s.students[id]
    if !exists {
        return errors.New("student not found")
    }

    delete(s.students, id)
    return nil
}