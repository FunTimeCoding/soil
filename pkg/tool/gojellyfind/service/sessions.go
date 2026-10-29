package service

import "github.com/funtimecoding/soil/pkg/jellyfin/session"

func (s *Service) Sessions() ([]*session.Session, error) {
	return s.source.Sessions()
}
