// OOP API
// รัน: cd api-patterns/oop ; go run .
//
// โครงสร้าง: จัดโฟลเดอร์ตาม "object / feature"
//
//	oop/
//	├── main.go
//	├── server/
//	│   └── server.go      object Server + interface Routable
//	└── student/            ทุกอย่างของนักเรียนอยู่ที่นี่
//	    ├── student.go     object Student (ข้อมูล + method)
//	    ├── store.go       object Store (ที่เก็บข้อมูล ห่อหุ้มไว้)
//	    └── handler.go     object Handler (รับ HTTP)
package main

import (
	"log"

	"oop-api/server"
	"oop-api/student"
)

func main() {
	store := student.NewStore()
	handler := student.NewHandler(store)

	srv := server.New(":8080")
	srv.Register(handler)

	log.Fatal(srv.Start())
}

