// SOLID API
// รัน: cd api-patterns/solid ; go run .
//
// โครงสร้าง: แยกโฟลเดอร์ตาม "หน้าที่เดียว" และมี interfaces/ เป็นศูนย์กลาง
//
//	solid/
//	├── main.go                                    ประกอบร่าง (Dependency Injection)
//	├── models/student.go                          รูปร่างข้อมูล
//	├── interfaces/student.go                      สัญญาทั้งหมด                       (I, D)
//	├── validators/student_validator.go            ตรวจข้อมูลอย่างเดียว                  (S)
//	├── repositories/
//	│   ├── memory_student_repository.go           เก็บข้อมูลอย่างเดียว                  (S)
//	│   └── logging_student_repository.go          เพิ่ม log โดยไม่แก้ของเดิม              (O, L)
//	├── services/student_service.go                business logic อย่างเดียว             (S, D)
//	└── handlers/student_handler.go                HTTP อย่างเดียว                       (S, D)
package main

import (
	"log"
	"net/http"

	"solid-api/handlers"
	"solid-api/repositories"
	"solid-api/services"
	"solid-api/validators"
)

func main() {
	// ลองลบบรรทัด Logging ออก ระบบก็ยังทำงานได้ (นี่คือ L และ O)
	repo := repositories.NewLoggingStudentRepository(
		repositories.NewMemoryStudentRepository(),
	)
	validator := validators.NewStudentValidator()
	service := services.NewStudentService(repo, validator)
	handler := handlers.NewStudentHandler(service)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	log.Println("SOLID API running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

