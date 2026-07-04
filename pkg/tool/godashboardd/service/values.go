package service

func (s *Service) Values() map[string][]string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	result := make(map[string][]string, len(s.values))

	for k, v := range s.values {
		result[k] = v
	}

	return result
}
