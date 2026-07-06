package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"net/http"
)

func (s *Server) Mount(m *http.ServeMux) {
	if s.authorization != nil {
		m.HandleFunc(
			fmt.Sprintf("GET %s", constant.SignInPath),
			s.signIn,
		)
		m.HandleFunc(
			fmt.Sprintf("GET %s", constant.CallbackPath),
			s.callback,
		)
		m.HandleFunc(
			fmt.Sprintf("GET %s", constant.SignOutPath),
			s.signOut,
		)
	}

	m.HandleFunc(
		"GET /palette",
		s.require(palette.NewServe(s.registry)),
	)
	m.HandleFunc("GET /{$}", s.require(s.dashboard))
	m.HandleFunc("GET /heatmap", s.require(s.heatmap))
	m.HandleFunc("POST /click", s.require(s.click))
	m.Handle("GET /event", s.require(s.event()))
	m.HandleFunc("GET /favicon.ico", s.favicon)
}
