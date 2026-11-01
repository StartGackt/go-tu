package repositories

import (
	"log"

	"solid-api/interfaces"
	"solid-api/models"
)

// LoggingStudentRepository "ห่อ" repository ตัวไหนก็ได้ แล้วเพิ่มการ log
//
// O (Open/Closed): เพิ่มความสามารถ (log) ได้ โดยไม่แก้โค้ด MemoryStudentRepository เลย
// L (Liskov Substitution): ใช้แทน repository ตัวจริงได้ทันที service ไม่รู้ตัวด้วยซ้ำ
type LoggingStudentRepository struct {
	next interfaces.StudentRepository
}

func NewLoggingStudentRepository(next interfaces.StudentRepository) *LoggingStudentRepository {
	return &LoggingStudentRepository{next: next}
}

func (r *LoggingStudentRepository) FindAll() []models.Student {
	log.Println("[repo] FindAll")
	return r.next.FindAll()
}

func (r *LoggingStudentRepository) FindByID(id int) (models.Student, error) {
	log.Println("[repo] FindByID", id)
	return r.next.FindByID(id)
}

func (r *LoggingStudentRepository) Create(s models.Student) models.Student {
	log.Println("[repo] Create", s.Name)
	return r.next.Create(s)
}

func (r *LoggingStudentRepository) Update(s models.Student) error {
	log.Println("[repo] Update", s.ID)
	return r.next.Update(s)
}

func (r *LoggingStudentRepository) Delete(id int) error {
	log.Println("[repo] Delete", id)
	return r.next.Delete(id)
}

