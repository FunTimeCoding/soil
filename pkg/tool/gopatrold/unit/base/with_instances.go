package base

import "github.com/funtimecoding/soil/pkg/patrol/mock_client"

func WithInstances(names ...string) Option {
	return func(s *Server) {
		s.names = names
		s.instances = map[string]*mock_client.Client{}

		for _, n := range names {
			s.instances[n] = mock_client.New()
		}
	}
}
