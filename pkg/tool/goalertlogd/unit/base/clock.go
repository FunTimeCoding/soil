package base

import "github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"

func (s *Server) Clock() *worker.Worker {
	s.t.Helper()

	for _, w := range s.Daemon().Workers() {
		if result, okay := w.(*worker.Worker); okay {
			return result
		}
	}

	s.t.Fatal("daemon hands out no clock")

	return nil
}
