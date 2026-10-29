package service

import "github.com/funtimecoding/soil/pkg/relational"

func (s *Service) Close() {
	if s.directoryGuard != nil {
		s.directoryGuard.Close()
	}

	s.authorization.Close()
	relational.CloseMapper(s.database)
}
