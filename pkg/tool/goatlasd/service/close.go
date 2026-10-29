package service

import "github.com/funtimecoding/soil/pkg/relational"

func (s *Service) Close() {
	s.netboxGuard.Close()

	for _, g := range s.guards {
		g.Close()
	}

	s.netbox.Close()

	for _, p := range s.peers {
		p.Close()
	}

	s.authorization.Close()
	relational.CloseMapper(s.database)
}
