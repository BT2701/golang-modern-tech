package week5

type StudentService struct {
	repo *StudentRepository
}

func NewStudentService(repo *StudentRepository) *StudentService {
	return &StudentService{repo: repo}
}

func (s *StudentService) GetAll() ([]Student, error) {
	return s.repo.GetAll()
}

func (s *StudentService) GetByID(id int) (*Student, error) {
	return s.repo.GetByID(id)
}

func (s *StudentService) Add(student *Student) (*Student, error) {
	err := s.repo.Add(student)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (s *StudentService) Update(student *Student) (*Student, error) {
	err := s.repo.Update(student)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (s *StudentService) Delete(id int) error {
	return s.repo.Delete(id)
}
