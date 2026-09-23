package service

func (s *Service) MarkSeen(session string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.store.MarkSeen(session)
}
