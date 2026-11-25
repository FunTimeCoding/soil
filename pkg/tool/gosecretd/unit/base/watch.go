package base

import "github.com/funtimecoding/soil/pkg/tool/gosecretd/watch"

func (s *Server) Watch() *watch.Watch {
	s.t.Helper()

	for _, w := range s.Daemon().Workers() {
		if result, okay := w.(*watch.Watch); okay {
			return result
		}
	}

	s.t.Fatal("daemon hands out no watch")

	return nil
}
