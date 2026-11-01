// Package http รวม route ของ adapter ฝั่ง HTTP
package http

import (
	"net/http"

	"clean-api/internal/adapter/http/handler"
)

func NewRouter(student *handler.StudentHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /students", student.List)
	mux.HandleFunc("GET /students/{id}", student.Get)
	mux.HandleFunc("POST /students", student.Create)
	mux.HandleFunc("PUT /students/{id}", student.Update)
	mux.HandleFunc("DELETE /students/{id}", student.Delete)
	return mux
}

