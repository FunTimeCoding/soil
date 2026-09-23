package service

func (s *Service) Stop() {
	s.doneOnce.Do(func() { close(s.done) })
}
