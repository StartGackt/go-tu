// Package controllers คือ C ใน MVC
// Controller = คนกลาง: รับ request -> สั่ง Model -> ส่งผลให้ View แสดง
package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"mvc-api/models"
	"mvc-api/views"
)

type StudentController struct {
	model *models.StudentModel
}

func NewStudentController(m *models.StudentModel) *StudentController {
	return &StudentController{model: m}
}


type studentRequest struct {
	Name  string `json:"names"`
	Score int    `json:"scores"`
}

// Index  GET /students
func (c *StudentController) Index(w http.ResponseWriter, r *http.Request) {
	views.RenderStudents(w, c.model.All())
}

// Show  GET /students/{id}
func (c *StudentController) Show(w http.ResponseWriter, r *http.Request) {
	s, err := c.model.Find(pathID(r))
	if err != nil {
		renderErr(w, err)
		return
	}
	views.RenderStudent(w, http.StatusOK, s)
}

// Store  POST /students
func (c *StudentController) Store(w http.ResponseWriter, r *http.Request) {
	var req studentRequest
	json.NewDecoder(r.Body).Decode(&req)
	s, err := c.model.Create(req.Name, req.Score)
	if err != nil {
		renderErr(w, err)
		return
	}
	views.RenderStudent(w, http.StatusCreated, s)
}

// Update  PUT /students/{id}
func (c *StudentController) Update(w http.ResponseWriter, r *http.Request) {
	var req studentRequest
	json.NewDecoder(r.Body).Decode(&req)
	s, err := c.model.Update(pathID(r), req.Name, req.Score)
	if err != nil {
		renderErr(w, err)
		return
	}
	views.RenderStudent(w, http.StatusOK, s)
}

// Destroy  DELETE /students/{id}
func (c *StudentController) Destroy(w http.ResponseWriter, r *http.Request) {
	if err := c.model.Delete(pathID(r)); err != nil {
		renderErr(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func pathID(r *http.Request) int {
	id, _ := strconv.Atoi(r.PathValue("id"))
	return id
}

func renderErr(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, models.ErrNotFound):
		views.RenderError(w, http.StatusNotFound, err.Error())
	case errors.Is(err, models.ErrInvalid):
		views.RenderError(w, http.StatusBadRequest, err.Error())
	default:
		views.RenderError(w, http.StatusInternalServerError, "internal error")
	}
}

