package service

import "github.com/funtimecoding/soil/pkg/chromium/tab"

func (s *Service) Tabs() ([]*tab.Tab, error) {
	return guarded(s.guard, s.source.Tabs)
}
