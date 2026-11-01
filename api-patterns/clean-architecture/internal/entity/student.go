// Package entity คือชั้นในสุด (Enterprise Business Rules)
// กฎ: ไม่ import อะไรในโปรเจกต์เลย และไม่มี json tag (ไม่รู้จัก HTTP)
package entity

import "errors"

var (
	ErrNotFound = errors.New("student not found")
	ErrInvalid  = errors.New("name is required and score must be 0-100")
)

type Student struct {
	ID    int
	Name  string
	Score int
}

func NewStudent(name string, score int) (*Student, error) {
	s := &Student{Name: name, Score: score}
	if err := s.Validate(); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Student) Validate() error {
	if s.Name == "" || s.Score < 0 || s.Score > 100 {
		return ErrInvalid
	}
	return nil
}

func (s *Student) Grade() string {
	switch {
	case s.Score >= 80:
		return "A"
	case s.Score >= 70:
		return "B"
	case s.Score >= 60:
		return "C"
	default:
		return "F"
	}
}

