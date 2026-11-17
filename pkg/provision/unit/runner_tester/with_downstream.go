package runner_tester

import "github.com/funtimecoding/soil/pkg/provision/face"

func WithDownstream(d face.Downstream) Option {
	return func(s *setup) {
		s.downstream = d
	}
}
