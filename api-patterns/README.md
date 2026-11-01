# API 4 รูปแบบ: OOP / MVC / SOLID / Clean Architecture

ทั้ง 4 โฟลเดอร์ทำ **API เดียวกันทุกอย่าง** (ระบบนักเรียน CRUD)
ต่างกันแค่ **วิธีวางโครงสร้างโฟลเดอร์** เปิดเทียบกันได้เลย
แต่ละโฟลเดอร์เป็นโปรเจกต์แยก (มี `go.mod` ของตัวเอง) ก๊อปไปใช้จริงได้ทันที

## API ที่ทุกโปรเจกต์มีเหมือนกัน
| Method | URL | ทำอะไร |
|---|---|---|
| GET | `/students` | ดูนักเรียนทั้งหมด |
| GET | `/students/{id}` | ดูนักเรียน 1 คน |
| POST | `/students` | เพิ่มนักเรียน |
| PUT | `/students/{id}` | แก้ไขนักเรียน |
| DELETE | `/students/{id}` | ลบนักเรียน |

---

## 1. OOP  -  จัดตาม "object / feature"
```
oop/
├── go.mod
├── main.go
├── server/
│   └── server.go          object Server + interface Routable
└── student/                ทุกอย่างของนักเรียนรวมอยู่ที่นี่
    ├── student.go         object Student (ข้อมูล + method)
    ├── store.go           object Store (ที่เก็บข้อมูล)
    └── handler.go         object Handler (รับ HTTP)
```
รัน: `cd api-patterns/oop ; go run .`

**จุดสังเกต:** เพิ่ม feature ใหม่ (เช่น ครู) = เพิ่มโฟลเดอร์ `teacher/` อีกโฟลเดอร์

---

## 2. MVC  -  จัดตาม "หน้าที่ M / V / C"
```
mvc/
├── go.mod
├── main.go
├── models/
│   └── student.go             M: ข้อมูล + กฎ + อ่าน/เขียนข้อมูล
├── views/
│   └── student_view.go        V: รูปแบบ JSON ที่ตอบกลับ
├── controllers/
│   └── student_controller.go  C: Index / Show / Store / Update / Destroy
└── routes/
    └── routes.go              จับคู่ URL กับ Controller
```
รัน: `cd api-patterns/mvc ; go run .`

**จุดสังเกต:** เพิ่ม feature ใหม่ = เพิ่มไฟล์ในแต่ละโฟลเดอร์ `models/teacher.go`, `views/teacher_view.go`, ...

---

## 3. SOLID  -  จัดตาม "หนึ่งโฟลเดอร์ หนึ่งหน้าที่" + มี interfaces เป็นศูนย์กลาง
```
solid/
├── go.mod
├── main.go
├── models/student.go                         รูปร่างข้อมูล
├── interfaces/student.go                     สัญญาทั้งหมด                    I, D
├── validators/student_validator.go           ตรวจข้อมูลอย่างเดียว               S
├── repositories/
│   ├── memory_student_repository.go          เก็บข้อมูลอย่างเดียว               S
│   └── logging_student_repository.go         เพิ่ม log โดยไม่แก้ของเดิม           O, L
├── services/student_service.go               business logic                  S, D
└── handlers/student_handler.go               HTTP อย่างเดียว                    S, D
```
รัน: `cd api-patterns/solid ; go run .`

**จุดสังเกต:** ลองยิง API แล้วดู log `[repo] ...` ใน terminal นั่นคือ OCP ทำงาน

---

## 4. Clean Architecture  -  จัดเป็น "ชั้น" แบบหัวหอม
```
clean-architecture/
├── go.mod
├── cmd/api/main.go                       จุดเริ่มโปรแกรม
└── internal/
    ├── entity/student.go                 ชั้น 1 (ในสุด) ไม่รู้จักใครเลย
    ├── usecase/
    │   ├── ports.go                      สัญญา Input/Output Port
    │   └── student_usecase.go            ชั้น 2
    └── adapter/                          ชั้น 3 (นอกสุด)
        ├── repository/memory/
        │   └── student_repository.go
        └── http/
            ├── router.go
            ├── dto/student_dto.go
            └── handler/student_handler.go
```
รัน: `cd api-patterns/clean-architecture ; go run ./cmd/api`

**จุดสังเกต:** `entity` ไม่มี `json` tag เลย เพราะชั้นในต้องไม่รู้จัก HTTP

---

## เทียบกัน
| | OOP | MVC | SOLID | Clean Architecture |
|---|---|---|---|---|
| จัดโฟลเดอร์ตาม | feature | M / V / C | หน้าที่ | ชั้น (layer) |
| จำนวนชั้น | 1-2 | 3 | 4-5 | 3 ชั้น + port |
| เหมาะกับ | โปรเจกต์เล็ก | เว็บทั่วไป | ทีมกลาง | โปรเจกต์ใหญ่ อยู่นาน |
| เปลี่ยน database | แก้ `store.go` | แก้ `models/` | เพิ่มไฟล์ใน `repositories/` | เพิ่มโฟลเดอร์ใน `adapter/repository/` |

## ทดสอบ (PowerShell) ใช้ได้กับทุกโปรเจกต์
```powershell
Invoke-RestMethod -Method Post http://localhost:8080/students -ContentType 'application/json' -Body '{"name":"Manee","score":85}'
Invoke-RestMethod http://localhost:8080/students
Invoke-RestMethod http://localhost:8080/students/1
Invoke-RestMethod -Method Put http://localhost:8080/students/1 -ContentType 'application/json' -Body '{"name":"Manee","score":95}'
Invoke-RestMethod -Method Delete http://localhost:8080/students/1
```

