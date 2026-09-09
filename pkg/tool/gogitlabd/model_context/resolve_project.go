package model_context

func (s *Server) resolveProject(identifier string) (int64, error) {
	return s.client.ResolveProject(identifier)
}
