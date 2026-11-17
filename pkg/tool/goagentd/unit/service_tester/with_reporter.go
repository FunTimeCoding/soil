package service_tester

import "github.com/funtimecoding/soil/pkg/face"

func WithReporter(r face.Reporter) Option {
	return func(s *setup) {
		s.reporter = r
	}
}
