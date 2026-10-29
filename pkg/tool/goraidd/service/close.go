package service

import "github.com/funtimecoding/soil/pkg/relational"

func (s *Service) Close() {
	s.closeOnce.Do(
		func() {
			s.parserGuard.Close()
			s.parser.Close()
			s.authorization.Close()
			relational.CloseMapper(s.database)
		},
	)
}
