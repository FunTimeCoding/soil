package model_context

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context/server"
	"net/http"
)

func (s *Server) Mount(m *http.ServeMux) {
	server.New(s.server).Setup(m)
}
