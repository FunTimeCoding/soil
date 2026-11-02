package service

import "github.com/funtimecoding/soil/pkg/relational"

func (s *Service) Close() {
	s.runner.Close()
	<-s.runner.Finished()

	for _, g := range s.guards {
		g.Close()
	}

	for _, c := range s.downstream {
		c.Close()
	}

	s.reaper.Close()
	relational.CloseMapper(s.database)
}
