// Package interfaces รวม "สัญญา" ทั้งหมดของระบบ
//
// I (Interface Segregation): แยก Reader / Writer ให้เล็ก
//
//	ใครแค่อ่าน ก็ขอแค่ StudentReader ไม่ต้องรู้จักการเขียน
//
// D (Dependency Inversion): ทุกชั้นพึ่ง interface ในไฟล์นี้ ไม่พึ่งของจริง
package interfaces

import "solid-api/models"

type StudentReader interface {
	FindAll() []models.Student
	FindByID(id int) (models.Student, error)
}

type StudentWriter interface {
	Create(s models.Student) models.Student
	Update(s models.Student) error
	Delete(id int) error
}

// StudentRepository = Reader + Writer
type StudentRepository interface {
	StudentReader
	StudentWriter
}

type StudentValidator interface {
	Validate(name string, score int) error
}

type StudentService interface {
	List() []models.Student
	Get(id int) (models.Student, error)
	Create(name string, score int) (models.Student, error)
	Update(id int, name string, score int) (models.Student, error)
	Delete(id int) error
}

