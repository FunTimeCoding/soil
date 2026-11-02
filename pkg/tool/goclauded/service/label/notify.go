package label

func (s *Service) notify() {
	s.notifier.Notify()
}
