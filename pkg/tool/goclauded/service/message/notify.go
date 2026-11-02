package message

func (s *Service) notify() {
	s.notifier.Notify()
}
