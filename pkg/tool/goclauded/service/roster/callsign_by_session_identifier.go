package roster

func (s *Service) CallsignBySessionIdentifier(identifier string) (string, error) {
	return s.store.CallsignBySessionIdentifier(identifier)
}
