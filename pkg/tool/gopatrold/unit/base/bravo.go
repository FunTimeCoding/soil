package base

import "github.com/funtimecoding/soil/pkg/patrol/mock_client"

func (s *Server) Bravo() *mock_client.Client {
	s.t.Helper()

	return s.Instance("bravo")
}
