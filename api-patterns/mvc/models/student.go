// Package models คือ M ใน MVC
// Model = ข้อมูล + กฎของข้อมูล + การอ่าน/เขียนข้อมูล
package models

import (
	"errors"
	"sort"
	"sync"
)

var (
	ErrNotFound = errors.New("student not found")
	ErrInvalid  = errors.New("name is required and score must be 0-100")
)

type Student struct {
	ID    int
	Name  string
	Score int
}

// Grade คือกฎของข้อมูล อยู่ใน Model
func (s Student) Grade() string {
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

func validate(name string, score int) error {
	if name == "" || score < 0 || score > 100 {
		return ErrInvalid
	}
	return nil
}

// StudentModel จัดการข้อมูลนักเรียนทั้งหมด
type StudentModel struct {
	mu     sync.RWMutex
	data   map[int]Student
	nextID int
}

func NewStudentModel() *StudentModel {
	return &StudentModel{data: map[int]Student{}, nextID: 1}
}

func (m *StudentModel) All() []Student {
	m.mu.RLock()
	defer m.mu.RUnlock()
	list := make([]Student, 0, len(m.data))
	for _, s := range m.data {
		list = append(list, s)
	}
	sort.Slice(list, func(i, j int) bool { return list[i].ID < list[j].ID })
	return list
}

func (m *StudentModel) Find(id int) (Student, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	s, ok := m.data[id]
	if !ok {
		return Student{}, ErrNotFound
	}
	return s, nil
}

func (m *StudentModel) Create(name string, score int) (Student, error) {
	if err := validate(name, score); err != nil {
		return Student{}, err
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	s := Student{ID: m.nextID, Name: name, Score: score}
	m.nextID++
	m.data[s.ID] = s
	return s, nil
}

func (m *StudentModel) Update(id int, name string, score int) (Student, error) {
	if err := validate(name, score); err != nil {
		return Student{}, err
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[id]; !ok {
		return Student{}, ErrNotFound
	}
	s := Student{ID: id, Name: name, Score: score}
	m.data[id] = s
	return s, nil
}

func (m *StudentModel) Delete(id int) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.data[id]; !ok {
		return ErrNotFound
	}
	delete(m.data, id)
	return nil
}

