package web

import (
	"github.com/funtimecoding/soil/pkg/web/palette"
	"net/http"
)

func (s *Server) Mount(m *http.ServeMux) {
	m.HandleFunc("GET /palette", palette.NewServe(s.registry))
	m.HandleFunc("GET /{$}", s.dashboard)
	m.HandleFunc("GET /search", s.searchPage)
	m.HandleFunc("GET /collections", s.collectionsPage)
	m.HandleFunc("GET /collections/{name}", s.documentsPage)
	m.HandleFunc("GET /documents/{path...}", s.documentDetailPage)
	m.HandleFunc("GET /favicon.ico", s.favicon)
}
