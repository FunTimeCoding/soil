package service

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Service) Close() {
	s.closeOnce.Do(
		func() {
			s.query.Close()
			errors.PanicClose(s.database)
		},
	)
}
