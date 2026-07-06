package web

import (
	"github.com/funtimecoding/soil/pkg/web/palette"
	"net/http"
)

func (s *Server) Mount(m *http.ServeMux) {
	m.HandleFunc("GET /palette", palette.NewServe(s.registry))
	m.HandleFunc("GET /palette/memories", s.paletteMemories)
	m.HandleFunc("GET /palette/memories/search", s.paletteMemoriesSearch)
	m.HandleFunc("GET /{$}", s.dashboard)
	m.HandleFunc("GET /memories", s.memoriesPage)
	m.HandleFunc("GET /memories/{identifier}", s.memoryDetailPage)
	m.HandleFunc("GET /impressions", s.impressionsPage)
	m.HandleFunc("GET /search", s.searchPage)
	m.HandleFunc("GET /favicon.ico", s.favicon)
}
