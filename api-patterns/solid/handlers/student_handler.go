// Package handlers มีหน้าที่เดียว: แปลง HTTP <-> service
package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"solid-api/interfaces"
	"solid-api/models"
)

type StudentHandler struct {
	service interfaces.StudentService // พึ่ง interface (D)
}

func NewStudentHandler(s interfaces.StudentService) *StudentHandler {
	return &StudentHandler{service: s}
}

func (h *StudentHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /students", h.List)
	mux.HandleFunc("GET /students/{id}", h.Get)
	mux.HandleFunc("POST /students", h.Create)
	mux.HandleFunc("PUT /students/{id}", h.Update)
	mux.HandleFunc("DELETE /students/{id}", h.Delete)
}

type studentRequest struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func (h *StudentHandler) List(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, h.service.List())
}

func (h *StudentHandler) Get(w http.ResponseWriter, r *http.Request) {
	s, err := h.service.Get(pathID(r))
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, s)
}

func (h *StudentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req studentRequest
	json.NewDecoder(r.Body).Decode(&req)
	s, err := h.service.Create(req.Name, req.Score)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, s)
}

func (h *StudentHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req studentRequest
	json.NewDecoder(r.Body).Decode(&req)
	s, err := h.service.Update(pathID(r), req.Name, req.Score)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, s)
}

func (h *StudentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if err := h.service.Delete(pathID(r)); err != nil {
		writeError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func pathID(r *http.Request) int {
	id, _ := strconv.Atoi(r.PathValue("id"))
	return id
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	switch {
	case errors.Is(err, models.ErrNotFound):
		status = http.StatusNotFound
	case errors.Is(err, models.ErrInvalid):
		status = http.StatusBadRequest
	}
	writeJSON(w, status, map[string]string{"error": err.Error()})
}

