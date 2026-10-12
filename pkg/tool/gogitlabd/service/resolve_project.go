package service

func (s *Service) ResolveProject(identifier string) (int64, error) {
	return s.client.ResolveProject(identifier)
}
