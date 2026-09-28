package service

func (s *Service) CleanExpiredLoginSessions() error {
	return s.store.CleanExpiredLoginSessions()
}
