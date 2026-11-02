package roster

func (s *Service) ResolveModelContextSession(
	modelContextSessionIdentifier string,
) (string, string, error) {
	return s.store.ResolveModelContextSession(modelContextSessionIdentifier)
}
