package service

func (s *Service) Subscribing() bool {
	return s.watcher != nil
}
