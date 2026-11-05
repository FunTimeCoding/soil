package service

import "github.com/funtimecoding/soil/pkg/iterm/session"

func (s *Service) CreateTab() (*session.Session, error) {
	return guarded(s.guard, s.client.CreateTab)
}
