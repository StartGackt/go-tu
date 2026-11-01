package student

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// Handler คือ object ที่รับ HTTP request ของนักเรียน
type Handler struct {
	store *Store
}

func NewHandler(store *Store) *Handler {
	return &Handler{store: store}
}

// RegisterRoutes ทำให้ Handler เป็นไปตาม interface server.Routable
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /students", h.list)
	mux.HandleFunc("GET /students/{id}", h.get)
	mux.HandleFunc("POST /students", h.create)
	mux.HandleFunc("PUT /students/{id}", h.update)
	mux.HandleFunc("DELETE /students/{id}", h.delete)
}

type request struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, h.store.All())
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	s, err := h.store.Get(pathID(r))
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, s)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var req request
	json.NewDecoder(r.Body).Decode(&req)
	s, err := New(req.Name, req.Score)
	if err != nil {
		writeError(w, err)
		return
	}
	h.store.Add(s)
	writeJSON(w, http.StatusCreated, s)
}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
	var req request
	json.NewDecoder(r.Body).Decode(&req)
	s, err := h.store.Update(pathID(r), req.Name, req.Score)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, s)
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	if err := h.store.Delete(pathID(r)); err != nil {
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
	case errors.Is(err, ErrNotFound):
		status = http.StatusNotFound
	case errors.Is(err, ErrInvalid):
		status = http.StatusBadRequest
	}
	writeJSON(w, status, map[string]string{"error": err.Error()})
}

