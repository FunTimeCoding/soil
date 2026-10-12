package server

func (s *Server) resolveInstance(instance *string) (string, error) {
	var name string

	if instance != nil {
		name = *instance
	}

	return s.service.ResolveInstance(name)
}
