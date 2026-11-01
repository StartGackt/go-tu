// Package services มีหน้าที่เดียว: business logic
package services

import (
	"solid-api/interfaces"
	"solid-api/models"
)

// StudentService รู้จักแค่ "สัญญา" (interface) ไม่รู้ว่าข้างหลังเป็น memory หรือ database
// D (Dependency Inversion): รับ dependency เข้ามาทาง constructor
type StudentService struct {
	repo      interfaces.StudentRepository
	validator interfaces.StudentValidator
}

func NewStudentService(repo interfaces.StudentRepository, v interfaces.StudentValidator) *StudentService {
	return &StudentService{repo: repo, validator: v}
}

func (s *StudentService) List() []models.Student {
	return s.repo.FindAll()
}

func (s *StudentService) Get(id int) (models.Student, error) {
	return s.repo.FindByID(id)
}

func (s *StudentService) Create(name string, score int) (models.Student, error) {
	if err := s.validator.Validate(name, score); err != nil {
		return models.Student{}, err
	}
	return s.repo.Create(models.Student{Name: name, Score: score}), nil
}

func (s *StudentService) Update(id int, name string, score int) (models.Student, error) {
	if err := s.validator.Validate(name, score); err != nil {
		return models.Student{}, err
	}
	st := models.Student{ID: id, Name: name, Score: score}
	if err := s.repo.Update(st); err != nil {
		return models.Student{}, err
	}
	return st, nil
}

func (s *StudentService) Delete(id int) error {
	return s.repo.Delete(id)
}

