package service

func (s *Service) closeBrowser() {
	s.browserMutex.Lock()
	defer s.browserMutex.Unlock()

	if s.browser != nil {
		s.browser.Close()
	}
}
