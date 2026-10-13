package model_context

import "context"

func (s *Server) resolveProject(
	x context.Context,
	identifier string,
) (string, int64, error) {
	instance, e := s.resolveInstance(x)

	if e != nil {
		return "", 0, e
	}

	project, f := s.service.ResolveProject(instance, identifier)

	if f != nil {
		return "", 0, f
	}

	return instance, project, nil
}
