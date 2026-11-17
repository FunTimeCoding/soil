package base

import "github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"

func WithReporter(r *memory.Memory) Option {
	return func(s *Server) {
		s.Reporter = r
	}
}
