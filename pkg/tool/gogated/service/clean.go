package service

func (s *Service) Clean() []error {
	var result []error

	if e := s.store.CleanExpiredLoginSessions(); e != nil {
		result = append(result, e)
	}

	if e := s.store.CleanExpiredAuthenticationSessions(); e != nil {
		result = append(result, e)
	}

	return result
}
