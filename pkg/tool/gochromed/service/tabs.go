package service

import "github.com/funtimecoding/soil/pkg/chromium/tab"

func (s *Service) Tabs() ([]*tab.Tab, error) {
	return s.source.Tabs()
}
