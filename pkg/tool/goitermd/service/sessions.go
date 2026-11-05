package service

import "github.com/funtimecoding/soil/pkg/iterm/session"

func (s *Service) Sessions() ([]*session.Session, error) {
	return guarded(s.guard, s.client.Sessions)
}
