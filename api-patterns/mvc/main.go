// MVC API
// รัน: cd api-patterns/mvc ; go run .
//
// โครงสร้าง: จัดโฟลเดอร์ตาม "หน้าที่" M / V / C
//
//	mvc/
//	├── main.go
//	├── models/
//	│   └── student.go              M: ข้อมูล + กฎ + อ่าน/เขียนข้อมูล
//	├── views/
//	│   └── student_view.go         V: รูปแบบ JSON ที่ตอบกลับ
//	├── controllers/
//	│   └── student_controller.go   C: รับ request สั่ง Model ส่งให้ View
//	└── routes/
//	    └── routes.go               จับคู่ URL กับ Controller
//
// ลำดับ: Request -> routes -> controllers -> models
//
//	-> views -> Response
package main

import (
	"log"
	"net/http"

	"mvc-api/controllers"
	"mvc-api/models"
	"mvc-api/routes"
)

func main() {
	studentModel := models.NewStudentModel()
	studentController := controllers.NewStudentController(studentModel)

	mux := http.NewServeMux()
	routes.Register(mux, studentController)

	log.Println("MVC API running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
