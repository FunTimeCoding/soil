package service

import "github.com/funtimecoding/soil/pkg/firefox/tab"

func (s *Service) Tabs() ([]*tab.Tab, error) {
	return s.client.Tabs()
}
