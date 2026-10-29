package service

func (s *Service) Notify() {
	s.notifier.Notify()
}
