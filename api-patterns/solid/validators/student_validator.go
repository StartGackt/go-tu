// Package validators มีหน้าที่เดียว: ตรวจข้อมูล
// S (Single Responsibility): แยกกฎการตรวจออกมา แก้กฎที่นี่ที่เดียว
package validators

import "solid-api/models"

type StudentValidator struct{}

func NewStudentValidator() *StudentValidator { return &StudentValidator{} }

func (StudentValidator) Validate(name string, score int) error {
	if name == "" || score < 0 || score > 100 {
		return models.ErrInvalid
	}
	return nil
}

