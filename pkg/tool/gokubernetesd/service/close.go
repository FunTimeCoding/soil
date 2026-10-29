package service

import "github.com/funtimecoding/soil/pkg/relational"

func (s *Service) Close() {
	for _, f := range s.PortForwards() {
		f.Stop()
	}

	for _, g := range s.guards {
		g.Close()
	}

	relational.CloseMapper(s.database)
}
