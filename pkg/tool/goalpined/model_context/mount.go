package model_context

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context/server"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func (s *Server) Mount(m *http.ServeMux) {
	server.New(s.server).Setup(guard.New(m, web.ServiceTokens()))
}
