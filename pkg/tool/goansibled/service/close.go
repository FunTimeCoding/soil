package service

import "github.com/funtimecoding/soil/pkg/relational"

func (s *Service) Close() {
	s.runner.Close()
	<-s.runner.Finished()

	if s.inventory != nil {
		s.inventory.Close()
	}

	s.reaper.Close()
	relational.CloseMapper(s.database)
}
