package system

func (s *Service) forget(name string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.cache, name)
}
