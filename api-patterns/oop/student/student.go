// Package student รวมทุกอย่างของ "นักเรียน" ไว้ในที่เดียว
// แนวคิด OOP: ข้อมูล (field) และพฤติกรรม (method) อยู่ด้วยกันใน object
package student

import "errors"

var (
	ErrNotFound = errors.New("student not found")
	ErrInvalid  = errors.New("name is required and score must be 0-100")
)

// Student คือ object นักเรียน
type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Score int    `json:"score"`
}

// New คือ constructor: สร้าง object ที่ถูกต้องเสมอ
func New(name string, score int) (*Student, error) {
	s := &Student{Name: name, Score: score}
	if err := s.Validate(); err != nil {
		return nil, err
	}
	return s, nil
}

// Validate คือพฤติกรรมของ object: ตรวจตัวเองได้
func (s *Student) Validate() error {
	if s.Name == "" || s.Score < 0 || s.Score > 100 {
		return ErrInvalid
	}
	return nil
}

// Update แก้ข้อมูลตัวเอง โดยตรวจก่อนเสมอ
func (s *Student) Update(name string, score int) error {
	tmp := Student{Name: name, Score: score}
	if err := tmp.Validate(); err != nil {
		return err
	}
	s.Name, s.Score = name, score
	return nil
}

