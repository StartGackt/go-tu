// Package usecase คือ Application Business Rules
// รู้จักแค่ entity เท่านั้น
package usecase

import "clean-api/internal/entity"

type studentInteractor struct {
	repo StudentRepository
}

func NewStudentUsecase(repo StudentRepository) StudentUsecase {
	return &studentInteractor{repo: repo}
}

func (u *studentInteractor) List() ([]entity.Student, error) {
	return u.repo.FindAll()
}

func (u *studentInteractor) Get(id int) (*entity.Student, error) {
	return u.repo.FindByID(id)
}

func (u *studentInteractor) Create(name string, score int) (*entity.Student, error) {
	s, err := entity.NewStudent(name, score)
	if err != nil {
		return nil, err
	}
	if err := u.repo.Create(s); err != nil {
		return nil, err
	}
	return s, nil
}

func (u *studentInteractor) Update(id int, name string, score int) (*entity.Student, error) {
	s, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	s.Name, s.Score = name, score
	if err := s.Validate(); err != nil {
		return nil, err
	}
	if err := u.repo.Update(s); err != nil {
		return nil, err
	}
	return s, nil
}

func (u *studentInteractor) Delete(id int) error {
	return u.repo.Delete(id)
}

