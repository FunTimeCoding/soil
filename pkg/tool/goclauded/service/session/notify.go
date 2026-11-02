package session

func (s *Service) notify() {
	s.notifier.Notify()
}
