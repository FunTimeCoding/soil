package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
)

func (s *Server) Served() *model_context_server.Server {
	s.t.Helper()

	if s.served != nil {
		return s.served
	}

	s.served = model_context_server.Mount(s.t, s.Daemon(), memory.New())
	s.t.Cleanup(s.served.Close)

	return s.served
}
