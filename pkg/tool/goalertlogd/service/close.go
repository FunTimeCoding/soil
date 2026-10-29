package service

import "github.com/funtimecoding/soil/pkg/relational"

func (s *Service) Close() {
	s.closeOnce.Do(
		func() {
			s.guard.Close()
			s.alertmanager.Close()
			relational.CloseMapper(s.database)
		},
	)
}
