package student

import (
	"sort"
	"sync"
)

// Store คือ object ที่เก็บนักเรียน
// Encapsulation: field ทั้งหมดเป็นตัวพิมพ์เล็ก คนนอกต้องใช้ผ่าน method เท่านั้น
type Store struct {
	mu     sync.RWMutex
	data   map[int]*Student
	nextID int
}

func NewStore() *Store {
	return &Store{data: map[int]*Student{}, nextID: 1}
}

func (st *Store) All() []Student {
	st.mu.RLock()
	defer st.mu.RUnlock()
	list := make([]Student, 0, len(st.data))
	for _, s := range st.data {
		list = append(list, *s)
	}
	sort.Slice(list, func(i, j int) bool { return list[i].ID < list[j].ID })
	return list
}

func (st *Store) Get(id int) (*Student, error) {
	st.mu.RLock()
	defer st.mu.RUnlock()
	s, ok := st.data[id]
	if !ok {
		return nil, ErrNotFound
	}
	c := *s
	return &c, nil
}

func (st *Store) Add(s *Student) {
	st.mu.Lock()
	defer st.mu.Unlock()
	s.ID = st.nextID
	st.nextID++
	st.data[s.ID] = s
}

func (st *Store) Update(id int, name string, score int) (*Student, error) {
	st.mu.Lock()
	defer st.mu.Unlock()
	s, ok := st.data[id]
	if !ok {
		return nil, ErrNotFound
	}
	if err := s.Update(name, score); err != nil {
		return nil, err
	}
	c := *s
	return &c, nil
}

func (st *Store) Delete(id int) error {
	st.mu.Lock()
	defer st.mu.Unlock()
	if _, ok := st.data[id]; !ok {
		return ErrNotFound
	}
	delete(st.data, id)
	return nil
}
