package model_context

import "context"

func (s *Server) resolveInstance(x context.Context) (string, error) {
	return s.instances.Resolve(x)
}
