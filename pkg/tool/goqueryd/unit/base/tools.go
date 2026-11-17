package base

import "github.com/funtimecoding/soil/pkg/generative/model_context_client"

func (s *Server) Tools() *model_context_client.Client {
	s.t.Helper()

	if s.tools == nil {
		s.tools = s.Connect(s.t)
	}

	return s.tools
}
