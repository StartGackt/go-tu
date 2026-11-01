// Package memory คือ adapter ฝั่งฐานข้อมูล (ทำตาม usecase.StudentRepository)
// อยากใช้ PostgreSQL ก็สร้าง adapter/repository/postgres/ แล้วเปลี่ยนใน main.go บรรทัดเดียว
package memory

import (
	"sort"
	"sync"

	"clean-api/internal/entity"
)

type StudentRepository struct {
	mu     sync.RWMutex
	data   map[int]entity.Student
	nextID int
}

func NewStudentRepository() *StudentRepository {
	return &StudentRepository{data: map[int]entity.Student{}, nextID: 1}
}

func (r *StudentRepository) FindAll() ([]entity.Student, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	list := make([]entity.Student, 0, len(r.data))
	for _, s := range r.data {
		list = append(list, s)
	}
	sort.Slice(list, func(i, j int) bool { return list[i].ID < list[j].ID })
	return list, nil
}

func (r *StudentRepository) FindByID(id int) (*entity.Student, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	s, ok := r.data[id]
	if !ok {
		return nil, entity.ErrNotFound
	}
	return &s, nil
}

func (r *StudentRepository) Create(s *entity.Student) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	s.ID = r.nextID
	r.nextID++
	r.data[s.ID] = *s
	return nil
}

func (r *StudentRepository) Update(s *entity.Student) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.data[s.ID]; !ok {
		return entity.ErrNotFound
	}
	r.data[s.ID] = *s
	return nil
}

func (r *StudentRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.data[id]; !ok {
		return entity.ErrNotFound
	}
	delete(r.data, id)
	return nil
}

