package service

func (s *Service) CleanExpiredAuthenticationSessions() error {
	return s.store.CleanExpiredAuthenticationSessions()
}
