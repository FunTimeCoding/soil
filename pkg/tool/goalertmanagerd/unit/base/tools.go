package base

import "github.com/funtimecoding/soil/pkg/generative/model_context_client"

func (s *Server) Tools() *model_context_client.Client {
	s.t.Helper()

	return s.Served().Connect(s.t)
}
