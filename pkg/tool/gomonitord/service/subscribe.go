package service

func (s *Service) Subscribe() chan struct{} {
	return s.notifier.Subscribe()
}
