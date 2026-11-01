// Clean Architecture API
// รัน: cd api-patterns/clean-architecture ; go run ./cmd/api
//
// โครงสร้าง: จัดเป็น "ชั้น" แบบหัวหอม ชั้นนอกพึ่งชั้นใน ชั้นในไม่รู้จักชั้นนอก
//
//	clean-architecture/
//	├── cmd/api/main.go                       จุดเริ่มโปรแกรม ประกอบร่าง
//	└── internal/
//	    ├── entity/                           ชั้น 1 (ในสุด): ข้อมูล + กฎธุรกิจ
//	    │   └── student.go
//	    ├── usecase/                          ชั้น 2: สิ่งที่ระบบทำได้
//	    │   ├── ports.go                      สัญญา (Input/Output Port)
//	    │   └── student_usecase.go
//	    └── adapter/                          ชั้น 3 (นอกสุด): ต่อกับโลกภายนอก
//	        ├── repository/memory/
//	        │   └── student_repository.go     ฝั่งฐานข้อมูล
//	        └── http/
//	            ├── router.go
//	            ├── dto/student_dto.go        แปลง entity <-> JSON
//	            └── handler/student_handler.go
//
// ทิศทางการพึ่งพา: adapter -> usecase -> entity
package main

import (
	"log"
	"net/http"

	httpadapter "clean-api/internal/adapter/http"
	"clean-api/internal/adapter/http/handler"
	"clean-api/internal/adapter/repository/memory"
	"clean-api/internal/usecase"
)

func main() {
	repo := memory.NewStudentRepository()
	uc := usecase.NewStudentUsecase(repo)
	h := handler.NewStudentHandler(uc)
	router := httpadapter.NewRouter(h)

	log.Println("Clean Architecture API running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

