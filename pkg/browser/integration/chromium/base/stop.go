package base

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Stack) Stop() {
	if s.process == nil {
		return
	}

	errors.LogOnError(s.process.Kill())
	s.process = nil
}
