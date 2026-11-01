// Package repositories มีหน้าที่เดียว: เก็บข้อมูล
package repositories

import (
	"sort"
	"sync"

	"solid-api/models"
)

// MemoryStudentRepository ทำตามสัญญา interfaces.StudentRepository
type MemoryStudentRepository struct {
	mu     sync.RWMutex
	data   map[int]models.Student
	nextID int
}

func NewMemoryStudentRepository() *MemoryStudentRepository {
	return &MemoryStudentRepository{data: map[int]models.Student{}, nextID: 1}
}

func (r *MemoryStudentRepository) FindAll() []models.Student {
	r.mu.RLock()
	defer r.mu.RUnlock()
	list := make([]models.Student, 0, len(r.data))
	for _, s := range r.data {
		list = append(list, s)
	}
	sort.Slice(list, func(i, j int) bool { return list[i].ID < list[j].ID })
	return list
}

func (r *MemoryStudentRepository) FindByID(id int) (models.Student, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	s, ok := r.data[id]
	if !ok {
		return models.Student{}, models.ErrNotFound
	}
	return s, nil
}

func (r *MemoryStudentRepository) Create(s models.Student) models.Student {
	r.mu.Lock()
	defer r.mu.Unlock()
	s.ID = r.nextID
	r.nextID++
	r.data[s.ID] = s
	return s
}

func (r *MemoryStudentRepository) Update(s models.Student) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.data[s.ID]; !ok {
		return models.ErrNotFound
	}
	r.data[s.ID] = s
	return nil
}

func (r *MemoryStudentRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.data[id]; !ok {
		return models.ErrNotFound
	}
	delete(r.data, id)
	return nil
}

