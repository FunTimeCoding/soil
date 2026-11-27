package service

import "github.com/funtimecoding/soil/pkg/browser/firefox/tab"

func (s *Service) CreateTab(l string) (*tab.Tab, error) {
	return s.client.CreateTab(l)
}
