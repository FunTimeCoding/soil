package service

func (s *Service) Close() {
	s.doneOnce.Do(func() { close(s.done) })
	s.browserMutex.Lock()
	defer s.browserMutex.Unlock()

	if s.browser != nil {
		s.browser.Close()
	}
}
