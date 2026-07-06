package web

import (
	"github.com/funtimecoding/soil/pkg/web/palette"
	"net/http"
)

func (s *Server) Mount(m *http.ServeMux) {
	m.HandleFunc("GET /palette", palette.NewServe(s.registry))
	m.HandleFunc("GET /{$}", s.dashboard)
	m.Handle("GET /event", s.event())
	m.HandleFunc("POST /move-up", s.moveUp)
	m.HandleFunc("POST /move-down", s.moveDown)
	m.HandleFunc("POST /reorder", s.reorder)
	m.HandleFunc("GET /favicon.ico", s.favicon)
}
