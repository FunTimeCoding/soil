package service

func (s *Service) Unsubscribe(c chan struct{}) {
	s.notifier.Unsubscribe(c)
}
