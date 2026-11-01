// Package dto แปลง entity <-> JSON
// entity ไม่มี json tag ดังนั้นรูปแบบ JSON เปลี่ยนได้ที่นี่ โดยไม่กระทบชั้นใน
package dto

import "clean-api/internal/entity"

type StudentRequest struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

type StudentResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Score int    `json:"score"`
	Grade string `json:"grade"`
}

func ToStudentResponse(s *entity.Student) StudentResponse {
	return StudentResponse{ID: s.ID, Name: s.Name, Score: s.Score, Grade: s.Grade()}
}

func ToStudentResponses(list []entity.Student) []StudentResponse {
	out := make([]StudentResponse, 0, len(list))
	for i := range list {
		out = append(out, ToStudentResponse(&list[i]))
	}
	return out
}

