// Package handler คือ adapter ฝั่ง HTTP: แปลง request -> เรียก usecase -> ตอบ JSON
package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"clean-api/internal/adapter/http/dto"
	"clean-api/internal/entity"
	"clean-api/internal/usecase"
)

type StudentHandler struct {
	uc usecase.StudentUsecase
}

func NewStudentHandler(uc usecase.StudentUsecase) *StudentHandler {
	return &StudentHandler{uc: uc}
}

func (h *StudentHandler) List(w http.ResponseWriter, r *http.Request) {
	list, err := h.uc.List()
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, dto.ToStudentResponses(list))
}

func (h *StudentHandler) Get(w http.ResponseWriter, r *http.Request) {
	s, err := h.uc.Get(pathID(r))
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, dto.ToStudentResponse(s))
}

func (h *StudentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.StudentRequest
	json.NewDecoder(r.Body).Decode(&req)
	s, err := h.uc.Create(req.Name, req.Score)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, dto.ToStudentResponse(s))
}

func (h *StudentHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req dto.StudentRequest
	json.NewDecoder(r.Body).Decode(&req)
	s, err := h.uc.Update(pathID(r), req.Name, req.Score)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, dto.ToStudentResponse(s))
}

func (h *StudentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if err := h.uc.Delete(pathID(r)); err != nil {
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
	case errors.Is(err, entity.ErrNotFound):
		status = http.StatusNotFound
	case errors.Is(err, entity.ErrInvalid):
		status = http.StatusBadRequest
	}
	writeJSON(w, status, map[string]string{"error": err.Error()})
}

