// Package routes รวม URL ทั้งหมด จับคู่กับ Controller
package routes

import (
	"net/http"

	"mvc-api/controllers"
)

func Register(mux *http.ServeMux, student *controllers.StudentController) {
	mux.HandleFunc("GET /students", student.Index)
	mux.HandleFunc("GET /students/{id}", student.Show)
	mux.HandleFunc("POST /students", student.Store)
	mux.HandleFunc("PUT /students/{id}", student.Update)
	mux.HandleFunc("DELETE /students/{id}", student.Destroy)
}

