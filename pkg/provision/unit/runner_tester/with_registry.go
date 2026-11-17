package runner_tester

import "github.com/funtimecoding/soil/pkg/face"

func withRegistry(r face.ProcessRegistry) Option {
	return func(s *setup) {
		s.registry = r
	}
}
