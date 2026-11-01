// Package server คือ object Server ที่รวม route ของทุก feature
package server

import (
	"log"
	"net/http"
)

// Routable คือ interface (สัญญา): object ไหนลงทะเบียน route ได้ ก็เสียบเข้า Server ได้
// Polymorphism: วันหน้ามี teacher.Handler, course.Handler ก็เสียบได้เหมือนกัน
type Routable interface {
	RegisterRoutes(mux *http.ServeMux)
}

type Server struct {
	addr string
	mux  *http.ServeMux
}

func New(addr string) *Server {
	return &Server{addr: addr, mux: http.NewServeMux()}
}

func (s *Server) Register(r Routable) {
	r.RegisterRoutes(s.mux)
}

func (s *Server) Start() error {
	log.Printf("OOP API running on http://localhost%s", s.addr)
	return http.ListenAndServe(s.addr, s.mux)
}

