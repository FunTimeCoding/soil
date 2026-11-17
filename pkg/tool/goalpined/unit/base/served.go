package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/daemon"
)

func (s *Server) Served() *model_context_server.Server {
	s.t.Helper()

	if s.served == nil {
		s.served = model_context_server.Mount(
			s.t,
			daemon.New(s.Option),
			s.Reporter,
		)
	}

	return s.served
}
