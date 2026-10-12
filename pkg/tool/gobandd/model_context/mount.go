package model_context

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context/server"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func (s *Server) Mount(g *guard.Mux) {
	server.New(s.server).Setup(g)
}
