package service

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Service) Close() {
	s.closeOnce.Do(
		func() {
			s.guard.Close()
			s.query.Close()
			errors.PanicClose(s.database)
		},
	)
}
