package service

import "github.com/funtimecoding/soil/pkg/iterm/session"

func (s *Service) CreateTab() (*session.Session, error) {
	return s.source.CreateTab()
}
