// Package views คือ V ใน MVC
// สำหรับ API, View = "หน้าตาของ JSON" ที่ส่งกลับไปให้ผู้ใช้
// View ไม่คิดคำนวณเอง แค่จัดรูปแบบข้อมูลจาก Model
package views

import (
	"encoding/json"
	"net/http"

	"mvc-api/models"
)

type studentJSON struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Score int    `json:"score"`
	Grade string `json:"grade"`
}

func toJSON(s models.Student) studentJSON {
	return studentJSON{ID: s.ID, Name: s.Name, Score: s.Score, Grade: s.Grade()}
}

func RenderStudent(w http.ResponseWriter, status int, s models.Student) {
	render(w, status, map[string]any{"data": toJSON(s)})
}

func RenderStudents(w http.ResponseWriter, list []models.Student) {
	out := make([]studentJSON, 0, len(list))
	for _, s := range list {
		out = append(out, toJSON(s))
	}
	render(w, http.StatusOK, map[string]any{"data": out, "total": len(out)})
}

func RenderError(w http.ResponseWriter, status int, msg string) {
	render(w, status, map[string]any{"error": msg})
}

func render(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

