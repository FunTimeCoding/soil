package base

import "github.com/funtimecoding/soil/pkg/generative/model_context_server"

func (s *Server) Served() *model_context_server.Server {
	s.t.Helper()

	if s.served == nil {
		s.served = model_context_server.Mount(s.t, s.daemon, s.Reporter)
	}

	return s.served
}
