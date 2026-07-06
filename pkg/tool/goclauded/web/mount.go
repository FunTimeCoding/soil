package web

import (
	"github.com/funtimecoding/soil/pkg/web/palette"
	"net/http"
)

func (s *Server) Mount(m *http.ServeMux) {
	m.HandleFunc("GET /palette", palette.NewServe(s.registry))
	m.HandleFunc("GET /{$}", s.dashboard)
	m.Handle("GET /event", s.event())
	m.HandleFunc("GET /sessions", s.sessionsPage)
	m.HandleFunc("GET /sessions/{identifier}", s.sessionDetailPage)
	m.HandleFunc("GET /sessions/{identifier}/edit", s.sessionEditForm)
	m.HandleFunc("POST /sessions/{identifier}/edit", s.sessionEditSubmit)
	m.HandleFunc("POST /sessions/{identifier}/pulse", s.sessionPulseSubmit)
	m.HandleFunc("POST /sessions/{identifier}/delete", s.sessionDeleteAction)
	m.HandleFunc("GET /activity", s.activityPage)
	m.HandleFunc("GET /messages", s.messagesPage)
	m.HandleFunc("GET /history", s.historyPage)
	m.HandleFunc("GET /history/{identifier}/edit", s.historyEditForm)
	m.HandleFunc("POST /history/{identifier}/edit", s.historyEditSubmit)
	m.HandleFunc("GET /favicon.ico", s.favicon)
	s.conversations.Mount(m)
}
