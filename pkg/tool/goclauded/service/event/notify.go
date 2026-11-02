package event

func (s *Service) notify() {
	s.notifier.Notify()
}
