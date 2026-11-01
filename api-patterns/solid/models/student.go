// Package models เก็บแค่รูปร่างข้อมูล (S: หน้าที่เดียว)
package models

import "errors"

var (
	ErrNotFound = errors.New("student not found")
	ErrInvalid  = errors.New("name is required and score must be 0-100")
)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Score int    `json:"score"`
}

