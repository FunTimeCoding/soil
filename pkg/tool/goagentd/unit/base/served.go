package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func (s *Server) Served() *model_context_server.Server {
	s.t.Helper()

	if s.served != nil {
		return s.served
	}

	d := s.Daemon()
	s.served = model_context_server.New(
		s.t,
		func(g *guard.Mux) {
			d.Mount(memory.New(), mock_recorder.New(), g)
		},
	)
	s.t.Cleanup(s.served.Close)

	return s.served
}
