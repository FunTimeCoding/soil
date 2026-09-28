package service

func (s *Service) CreateLoginSession(authorizeRequest string) (string, error) {
	return s.store.CreateLoginSession(authorizeRequest)
}
