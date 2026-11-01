package usecase

import "clean-api/internal/entity"

// StudentRepository คือ "Output Port"
// usecase เป็นคนกำหนดสัญญา ชั้นนอก (adapter) ต้องทำตาม
// ลูกศรการพึ่งพาจึงชี้เข้าใน: adapter -> usecase
type StudentRepository interface {
	FindAll() ([]entity.Student, error)
	FindByID(id int) (*entity.Student, error)
	Create(s *entity.Student) error
	Update(s *entity.Student) error
	Delete(id int) error
}

// StudentUsecase คือ "Input Port" ที่ชั้นนอก (handler) เรียกใช้
type StudentUsecase interface {
	List() ([]entity.Student, error)
	Get(id int) (*entity.Student, error)
	Create(name string, score int) (*entity.Student, error)
	Update(id int, name string, score int) (*entity.Student, error)
	Delete(id int) error
}

