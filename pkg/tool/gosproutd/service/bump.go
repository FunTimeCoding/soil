package service

func (s *Service) Bump(identifier uint) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.store.Bump(identifier)
	s.notifier.Notify()
}
