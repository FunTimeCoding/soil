package base

import "github.com/funtimecoding/soil/pkg/worker"

func (s *Server) Tick() {
	s.t.Helper()

	for _, w := range s.daemon.Workers() {
		if c, okay := w.(*worker.Worker); okay {
			c.Poll()
		}
	}
}
