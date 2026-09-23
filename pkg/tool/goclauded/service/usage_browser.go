package service

import "github.com/funtimecoding/soil/pkg/chromium"

func (s *Service) usageBrowser() *chromium.Client {
	s.browserMutex.Lock()
	defer s.browserMutex.Unlock()

	if s.browser == nil {
		s.browser = chromium.NewEnvironment()
	}

	return s.browser
}
