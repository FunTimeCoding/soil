package base

import "github.com/funtimecoding/soil/pkg/patrol/mock_client"

func (s *Server) Instance(name string) *mock_client.Client {
	s.t.Helper()
	result, okay := s.instances[name]

	if !okay {
		s.t.Fatalf("no instance %s", name)
	}

	return result
}
