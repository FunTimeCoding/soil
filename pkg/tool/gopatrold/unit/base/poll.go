package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gopatrold/worker"
)

func (s *Server) Poll() {
	s.t.Helper()
	w := s.Daemon().Workers()
	clock, okay := w[len(w)-1].(*worker.Worker)
	assert.True(s.t, okay)
	clock.Poll()
}
