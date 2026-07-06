package web

import (
	"github.com/funtimecoding/soil/pkg/web/palette"
	"net/http"
)

func (s *Server) Mount(m *http.ServeMux) {
	m.HandleFunc("GET /palette", palette.NewServe(s.registry))
	m.HandleFunc("GET /{$}", s.dashboard)
	m.HandleFunc("GET /entries", s.entries)
	m.HandleFunc("GET /add", s.add)
	m.HandleFunc("POST /add", s.addSubmit)
	m.HandleFunc("GET /entry/detail", s.detail)
	m.HandleFunc("GET /entry/edit", s.edit)
	m.HandleFunc("POST /entry/edit", s.editSubmit)
	m.HandleFunc("POST /entry/delete", s.delete)
}
