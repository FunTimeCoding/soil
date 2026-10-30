package service

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Service) Close() {
	s.closeOnce.Do(
		func() {
			s.guard.Close()
			s.embedder.Close()
			errors.LogClose(s.reranker)
			errors.LogClose(s.counter)
			errors.PanicClose(s.database)
		},
	)
}
