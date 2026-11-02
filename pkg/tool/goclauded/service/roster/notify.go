package roster

func (s *Service) notify() {
	s.notifier.Notify()
}
